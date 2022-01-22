package crawl

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/url"
	"strings"

	"github.com/cnk3x/urlx"
	"github.com/goccy/go-json"
	"github.com/google/go-querystring/query"
)

// JSONBody 提交JSON
func JSONBody(in any) urlx.Body {
	return func() (contentType string, body io.Reader, err error) {
		contentType = "application/json; charset=utf-8"
		switch o := in.(type) {
		case io.Reader:
			body = o
		case []byte:
			body = bytes.NewReader(o)
		case string:
			body = strings.NewReader(o)
		case bytes.Buffer:
			body = bytes.NewReader(o.Bytes())
		case *bytes.Buffer:
			body = bytes.NewReader(o.Bytes())
		default:
			var data []byte
			if data, err = json.Marshal(in); err == nil {
				body = bytes.NewReader(data)
			}
		}
		return
	}
}

// FormBody 提交Form表单
func FormBody(in any) urlx.Body {
	return func() (contentType string, body io.Reader, err error) {
		contentType = "application/x-www-form-urlencoded; charset=utf-8"
		switch o := in.(type) {
		case io.Reader:
			body = o
		case []byte:
			body = bytes.NewReader(o)
		case string:
			body = strings.NewReader(o)
		case bytes.Buffer:
			body = &o
		case *bytes.Buffer:
			body = o
		case url.Values:
			body = strings.NewReader(o.Encode())
		case *url.Values:
			body = strings.NewReader(o.Encode())
		case map[string]string:
			values := url.Values{}
			for k, v := range o {
				values.Set(k, v)
			}
			body = strings.NewReader(values.Encode())
		default:
			if r, ok := o.(io.Reader); ok {
				body = r
			} else {
				var values url.Values
				if values, err = query.Values(in); err == nil {
					body = strings.NewReader(values.Encode())
				}
			}
		}
		return
	}
}

func XMLBody(in any) urlx.Body {
	return func() (contentType string, body io.Reader, err error) {
		contentType = "application/xml; charset=utf-8"
		switch o := in.(type) {
		case io.Reader:
			body = o
		case []byte:
			body = bytes.NewReader(o)
		case string:
			body = strings.NewReader(o)
		case bytes.Buffer:
			body = bytes.NewReader(o.Bytes())
		case *bytes.Buffer:
			body = bytes.NewReader(o.Bytes())
		default:
			var data []byte
			if data, err = xml.Marshal(in); err == nil {
				body = bytes.NewReader(data)
			}
		}
		return
	}
}
