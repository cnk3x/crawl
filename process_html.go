package crawl

import (
	"fmt"
	"io"
	"net/http"

	"github.com/cnk3x/urlx"
	"github.com/PuerkitoBio/goquery"
)

// ProcessHtml Html选择器
func ProcessHtml(readHtml func(doc *goquery.Selection) error) urlx.Process {
	return func(resp *http.Response, body io.ReadCloser) error {
		defer body.Close()
		if resp.StatusCode >= 400 {
			return fmt.Errorf("%s(%d)", resp.Status, resp.StatusCode)
		}
		doc, err := goquery.NewDocumentFromReader(body)
		if err != nil {
			return fmt.Errorf("read as html: %w", err)
		}
		return readHtml(doc.Selection)
	}
}

// ProcessStruct Html解析到Struct
func ProcessStruct(rootSelect string, options StrcutOptions, out any) urlx.Process {
	return ProcessHtml(func(doc *goquery.Selection) error {
		return BindStruct(doc.Find(rootSelect), out, options)
	})
}

// ProcessMap Html解析到Map, out must map[string]any or []map[string]any
func ProcessMap(rootSelect string, options MapOptions, out any) urlx.Process {
	return ProcessHtml(func(doc *goquery.Selection) error {
		data, err := BindMapField(doc, options.MapField)
		if err != nil {
			return err
		}
		if len(data) == 0 {
			return nil
		}

		var ok bool
		switch t := out.(type) {
		case *map[string]any:
			*t, ok = data[0].(map[string]any)
		case *[]map[string]any:
			*t = make([]map[string]any, len(data))
			for i, x := range data {
				if (*t)[i], ok = x.(map[string]any); !ok {
					break
				}
			}
		}
		if !ok {
			return fmt.Errorf("can not set %T to %T", data, out)
		}
		return nil
	})
}
