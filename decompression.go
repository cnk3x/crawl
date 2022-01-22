package crawl

import (
	"io"
	"net/http"
	"strings"

	"github.com/cnk3x/urlx"
	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/flate"
	"github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/s2"
	"github.com/klauspost/compress/snappy"
	"github.com/klauspost/compress/zstd"
)

// type (
// 	Process      = func(resp *http.Response, body io.ReadCloser) error // 响应处理器
// 	HeaderOption = func(headers http.Header)                           // 请求头处理
// )

const (
	HeaderAcceptEncoding  = "Accept-Encoding"
	HeaderContentEncoding = "Content-Encoding"
)

// Decompression 解压Body
func Decompression(next urlx.Process) urlx.Process {
	return func(resp *http.Response, body io.ReadCloser) (err error) {
		defer body.Close()
		contentEncoding := resp.Header.Get(HeaderContentEncoding)
		if contentEncoding != "" {
			decoded := true
			switch contentEncoding {
			case "br":
				body = io.NopCloser(brotli.NewReader(body))
			case "deflate":
				body = flate.NewReader(body)
			case "gzip":
				body, err = gzip.NewReader(body)
			case "s2":
				body = io.NopCloser(s2.NewReader(body))
			case "snappy":
				body = io.NopCloser(snappy.NewReader(body))
			case "zstd":
				b, er := zstd.NewReader(body)
				if er != nil {
					return er
				}
				body = b.IOReadCloser()
			default:
				decoded = false
			}
			if err != nil {
				return
			}
			if decoded {
				resp.Header.Del(HeaderContentEncoding)
			}
		}
		return next(resp, io.NopCloser(body))
	}
}

// AcceptEncoding 接受编码
func AcceptEncoding(acceptEncodings ...string) urlx.HeaderOption {
	return func(headers http.Header) {
		headers.Set(HeaderAcceptEncoding, strings.Join(acceptEncodings, ","))
	}
}

var (
	// AcceptAllEncodings 接受所有的编码格式
	AcceptAllEncodings = AcceptEncoding("zstd", "br", "gzip", "deflate", "snappy", "s2")
	// DefaultEncodings 默认接受所有的编码格式
	DefaultEncodings = AcceptEncoding("gzip", "deflate", "br")
)
