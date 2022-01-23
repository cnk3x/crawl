package crawl

import (
	"bytes"
	"github.com/cnk3x/urlx"
	"io"
	"net/http"
	"net/http/httputil"
)

var dumpBR = []byte{'\n'}
var dumpLine = append(bytes.Repeat([]byte("-"), 70), dumpBR...)

func Dump(w io.Writer, reqBody, respBody bool) urlx.ProcessMw {
	return func(next urlx.Process) urlx.Process {
		return func(resp *http.Response, body io.ReadCloser) error {
			w.Write(dumpLine)
			reqDump, err := httputil.DumpRequest(resp.Request, reqBody)
			if err != nil {
				w.Write([]byte(err.Error()))
			} else {
				w.Write(reqDump)
			}
			w.Write(dumpBR)
			w.Write(dumpLine)
			resDump, _ := httputil.DumpResponse(resp, respBody)
			if err != nil {
				w.Write([]byte(err.Error()))
			} else {
				w.Write(resDump)
			}
			w.Write(dumpBR)
			w.Write(dumpLine)
			return next(resp, body)
		}
	}
}
