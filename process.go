package crawl

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/cnk3x/urlx"
	"github.com/goccy/go-json"
	"github.com/goccy/go-yaml"
)

// ProcessJSON 处理JSON响应
func ProcessJSON(out any) urlx.Process {
	return func(resp *http.Response, body io.ReadCloser) error {
		defer body.Close()
		return json.NewDecoder(body).Decode(out)
	}
}

// ProcessYAML 处理yaml响应
func ProcessYAML(out any) urlx.Process {
	return func(resp *http.Response, body io.ReadCloser) error {
		defer body.Close()
		return yaml.NewDecoder(body).Decode(out)
	}
}

// ProcessXML 处理xml响应
func ProcessXML(out any) urlx.Process {
	return func(resp *http.Response, body io.ReadCloser) error {
		defer body.Close()
		return xml.NewDecoder(body).Decode(out)
	}
}
