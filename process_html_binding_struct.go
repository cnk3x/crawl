package crawl

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrValueCannotAddress = errors.New("value can not address")
	ErrValueNotBasicKind  = errors.New("value not a basic kind")
	ErrValueCast          = errors.New("value string cast error")
)

type StructOptions struct {
	SelectTag string // select
	AttrTag   string // attr
	FormatTag string // format
	FindTag   string // find
	ReplTag   string // repl
}

func BindStruct(sel *goquery.Selection, out any, options StructOptions) error {
	if sel.Size() > 0 {
		return bindStruct(reflect.ValueOf(out), sel, strcutTags{}, options)
	}
	return nil
}

func bindStruct(rv reflect.Value, sel *goquery.Selection, tags strcutTags, options StructOptions) error {
	if rv = reflect.Indirect(rv); !rv.CanAddr() {
		return ErrValueCannotAddress
	}
	rt := rv.Type()

	if isStructBasicType(rt) {
		err := setStructSelectText(rv, sel, tags)
		if err != nil {
			return fmt.Errorf("set basic type: %s: %w", rt, err)
		}
		return err
	}

	kind := rt.Kind()
	switch kind {
	case reflect.Struct:
		for i := 0; i < rt.NumField(); i++ {
			fs := rt.Field(i)
			if fs.Anonymous {
				if err := bindStruct(rv.Field(i), sel, tags, options); err != nil {
					return fmt.Errorf("set anonymous struct: %w", err)
				}
			}
			childTags := parseStructOptions(fs, options)
			if childTags.Select == "" || childTags.Select == "-" {
				continue
			}
			if err := bindStruct(rv.Field(i), sel.Find(childTags.Select), childTags, options); err != nil {
				return fmt.Errorf("set struct field: %s: %w", fs.Name, err)
			}
		}
	case reflect.Slice:
		it := rt.Elem()
		var ip bool
		if ip = it.Kind() == reflect.Ptr; ip {
			it = it.Elem()
		}
		var err error
		sel.EachWithBreak(func(_ int, el *goquery.Selection) bool {
			iv := reflect.New(it).Elem()
			if err = bindStruct(iv, el, tags, options); err == nil {
				if ip {
					iv = iv.Addr()
				}
				rv.Set(reflect.Append(rv, iv))
			}
			return err == nil
		})
		if err != nil {
			return fmt.Errorf("set slice: %w", err)
		}
	default:
		return fmt.Errorf("%w: %s", ErrValueCannotAddress, kind)
	}
	return nil
}

func setStructSelectText(fv reflect.Value, sel *goquery.Selection, tags strcutTags) error {
	s := getStructSelectionText(sel, tags.Attr, tags.Find, tags.Repl)
	if s != "" {
		return setStructBasicValue(fv, s, tags.Format)
	}
	return nil
}

func setStructBasicValue(fv reflect.Value, value, format string) error {
	if fv = reflect.Indirect(fv); !fv.CanAddr() {
		return ErrValueCannotAddress
	}

	if value == "" {
		return nil
	}

	switch fv.Type() {
	case durationType:
		if d, err := time.ParseDuration(value); err != nil {
			return fmt.Errorf("%w: %s", ErrValueCast, err.Error())
		} else {
			fv.SetInt(int64(d))
		}
		return nil
	case durationType2:
		if d, err := ParseDuration(value); err != nil {
			return fmt.Errorf("%w: %s", ErrValueCast, err.Error())
		} else {
			fv.SetInt(int64(d))
		}
		return nil
	case timeType:
		if format == "" {
			format = time.RFC3339
		}
		if t, err := time.Parse(format, value); err != nil {
			return fmt.Errorf("%w: %s", ErrValueCast, err.Error())
		} else {
			fv.Set(reflect.ValueOf(t))
		}
		return nil
	case bytesType:
		fv.Set(reflect.ValueOf([]byte(value)))
		return nil
	default:
		switch fv.Kind() {
		case reflect.String:
			fv.SetString(value)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(value, 0, 0)
			if err != nil {
				return fmt.Errorf("%w: %s", ErrValueCast, err.Error())
			}
			fv.SetInt(x)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(value, 0, 0)
			if err != nil {
				return fmt.Errorf("%w: %s", ErrValueCast, err.Error())
			}
			fv.SetUint(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(value, 0)
			if err != nil {
				return fmt.Errorf("%w: %s", ErrValueCast, err.Error())
			}
			fv.SetFloat(x)
		case reflect.Bool:
			x, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("%w: %s", ErrValueCast, err.Error())
			}
			fv.SetBool(x)
		default:
			return fmt.Errorf("%w: %s(%s): %s", ErrValueNotBasicKind, fv.Type(), fv.Kind(), value)
		}
	}

	return nil
}

func isStructBasicType(ft reflect.Type) bool {
	switch ft {
	case durationType, timeType, bytesType:
		return true
	default:
		switch ft.Kind() {
		case reflect.String:
		case reflect.Int64:
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		case reflect.Float32, reflect.Float64:
		case reflect.Bool:
		default:
			return false
		}
		return true
	}
}

func getStructSelectionText(sel *goquery.Selection, attr, find, repl string) (s string) {
	sel = sel.First()
	switch attr {
	case "", "text":
		s = sel.Text()
	case "html":
		s, _ = sel.Html()
	default:
		s = sel.AttrOr(attr, "")
	}
	if s = strings.TrimSpace(s); s != "" && find != "" {
		if re, _ := regexp.Compile(find); re != nil {
			if s = re.FindString(s); s != "" && repl != "" {
				s = re.ReplaceAllString(s, repl)
			}
		}
	}
	return s
}

func parseStructOptions(fs reflect.StructField, options StructOptions) (tags strcutTags) {
	if options.SelectTag == "" {
		options.SelectTag = "select"
	}
	if options.AttrTag == "" {
		options.AttrTag = "attr"
	}
	if options.FindTag == "" {
		options.FindTag = "find"
	}
	if options.ReplTag == "" {
		options.ReplTag = "repl"
	}
	if options.FormatTag == "" {
		options.FormatTag = "format"
	}

	tags.Select = fs.Tag.Get(options.SelectTag)
	tags.Find = fs.Tag.Get(options.FindTag)
	tags.Repl = fs.Tag.Get(options.ReplTag)
	tags.Format = fs.Tag.Get(options.FormatTag)
	if tags.Attr = fs.Tag.Get(options.AttrTag); tags.Attr == "" {
		tags.Attr = "text"
	}
	return
}

var (
	timeType      = reflect.TypeOf(time.Time{})
	durationType  = reflect.TypeOf(time.Duration(0))
	durationType2 = reflect.TypeOf(Duration(0))
	bytesType     = reflect.TypeOf([]byte{})
)

type strcutTags struct {
	Select string
	Attr   string
	Format string
	Find   string
	Repl   string
}
