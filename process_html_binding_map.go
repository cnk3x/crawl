package crawl

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type MapField struct {
	Name   string     `json:"name,omitempty" xml:"name,omitempty"`     // 字段名称
	Value  string     `json:"value,omitempty" xml:"value,omitempty"`   // 模板值
	Select string     `json:"select,omitempty" xml:"select,omitempty"` // 选择器
	Attr   string     `json:"attr,omitempty" xml:"attr,omitempty"`     // 属性选择
	Format string     `json:"format,omitempty" xml:"format,omitempty"` // 格式化
	Find   string     `json:"find,omitempty" xml:"find,omitempty"`     // 结果再查找（正则表达式）
	Repl   string     `json:"repl,omitempty" xml:"repl,omitempty"`     // 结果查找后再替换（正则替换表达式）
	List   bool       `json:"list,omitempty" xml:"list,omitempty"`     // 是否列表
	Split  string     `json:"split,omitempty" xml:"split,omitempty"`   // 是否对字段再进行拆分
	Type   string     `json:"type,omitempty" xml:"type,omitempty"`     // 类型: time, duration, string, int, float, bool, 默认 string
	Fields []MapField `json:"fields,omitempty" xml:"fields,omitempty"` // 字段
}

func BindMapField(doc *goquery.Selection, params map[string]string, field MapField) ([]any, error) {
	out, err := bindMapField(doc, params, field, false)
	if out != nil {
		slice, ok := out.([]any)
		if !ok {
			slice = []any{out}
		}
		return slice, err
	}
	return nil, err
}

func bindMapField(doc *goquery.Selection, params map[string]string, field MapField, iter bool) (any, error) {
	if field.Select != "" && !iter {
		doc = doc.Find(field.Select)
	}

	var s string
	if field.Value == "" {
		if !iter && field.List {
			var out []any
			var err error
			doc.EachWithBreak(func(_ int, itemDoc *goquery.Selection) bool {
				if arrayItem, ex := bindMapField(itemDoc, params, field, true); ex == nil {
					if s, ok := arrayItem.(string); ok {
						if s != "" {
							if field.Split != "" && field.Split != "false" {
								if field.Split == "true" {
									field.Split = `\s+`
								}
								if re, _ := regexp.Compile(field.Split); re != nil {
									ss := regexp.MustCompile(field.Split).Split(s, -1)
									anyT := make([]any, len(ss))
									for i, v := range ss {
										anyT[i] = strings.TrimSpace(v)
									}
									out = append(out, anyT...)
								} else {
									out = append(out, s)
								}
							} else {
								out = append(out, s)
							}
						}
					} else {
						out = append(out, arrayItem)
					}
				} else {
					err = ex
				}
				return err == nil
			})
			return out, err
		}

		doc = doc.First()

		// 有子节点
		if len(field.Fields) > 0 {
			out := map[string]any{}
			var err error
			for _, child := range field.Fields {
				if fieldItem, ex := bindMapField(doc, params, child, false); ex == nil {
					out[child.Name] = fieldItem
				} else {
					err = ex
					break
				}
			}
			return out, err
		}

		switch field.Attr {
		case "", "text":
			s = doc.Text()
		case "html":
			s, _ = doc.Html()
		default:
			s, _ = doc.Attr(field.Attr)
		}
	} else {
		s = ReplaceTemplate(field.Value, params)
	}

	if s = strings.TrimSpace(s); s != "" {
		if field.Find != "" {
			if re, _ := regexp.Compile(field.Find); re != nil {
				if s = re.FindString(s); s != "" {
					if field.Repl != "" {
						s = re.ReplaceAllString(s, field.Repl)
					}
					s = strings.TrimSpace(s)
				}
			}
		}

		var v any
		switch field.Type {
		case "time":
			if field.Format == "" {
				field.Format = time.RFC3339
			}
			v, _ = time.Parse(field.Format, s)
		case "duration":
			v, _ = ParseDuration(s)
		case "int":
			v, _ = strconv.ParseInt(s, 0, 0)
		case "float":
			v, _ = strconv.ParseFloat(s, 0)
		case "bool":
			v, _ = strconv.ParseBool(s)
		default:
			v = s
		}
		return v, nil
	}

	return "", nil
}
