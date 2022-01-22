package crawl

import (
	"io"
	"mime"
	"net/http"
	"strings"

	"github.com/cnk3x/urlx"
	"golang.org/x/text/encoding/htmlindex"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const (
	HeaderContentType = "Content-Type"
	ParamCharset      = "charset"
)

// Charset 指定响应的编码，auto 或者空则通过 Content-Type 自动判断
func Charset(charset string) urlx.ProcessMw {
	getCharset := func(mimeType string, params map[string]string) string {
		if charset == "" || charset == "auto" {
			if strings.HasPrefix(mimeType, "text/") {
				if len(params) > 0 {
					if charset = strings.TrimSpace(params[ParamCharset]); charset != "" {
						return charset
					}
				}
			}
		}
		return strings.ToLower(charset)
	}

	return func(next urlx.Process) urlx.Process {
		return func(resp *http.Response, body io.ReadCloser) error {
			defer body.Close()
			var r io.Reader = body
			mimeType, params, _ := mime.ParseMediaType(resp.Header.Get(HeaderContentType))
			if charset := getCharset(mimeType, params); charset != "" && charset != "UTF-8" {
				if codec, err := htmlindex.Get(charset); err == nil && codec != unicode.UTF8 {
					r = transform.NewReader(r, codec.NewDecoder())
					resp.Header.Set(HeaderContentType, mimeType)
				}
			}
			return next(resp, io.NopCloser(r))
		}
	}
}

// AutoCharset 将响应解码成UTF-8
var AutoCharset = Charset("auto")
