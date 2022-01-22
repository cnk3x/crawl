package crawl

import (
	"fmt"
	"io"
	"time"

	"github.com/valyala/fasttemplate"
)

// ReplaceTemplate 模板替换
func ReplaceTemplate(template string, args map[string]any) (s string) {
	s, _ = fasttemplate.ExecuteFuncStringWithErr(template, "{", "}", func(w io.Writer, tag string) (int, error) {
		v := args[tag]
		switch t := v.(type) {
		case time.Duration:
			return w.Write([]byte(t.String()))
		case string:
			return w.Write([]byte(t))
		case []byte:
			return w.Write(t)
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			return fmt.Fprintf(w, "%d", t)
		case float32, float64:
			return fmt.Fprintf(w, "%f", t)
		case bool:
			if t {
				return w.Write([]byte("true"))
			}
			return w.Write([]byte("false"))
		case time.Time:
			return w.Write([]byte(t.Format(time.RFC3339)))
		default:
			if t, ok := v.(fmt.Stringer); ok {
				return w.Write([]byte(t.String()))
			}
			if t, ok := v.(error); ok {
				return w.Write([]byte(t.Error()))
			}
			return 0, nil
		}
	})
	return
}

// ReplaceTemplate 模板替换
func ReplaceTemplateWithFlatArgs(template string, args ...any) (s string) {
	return ReplaceTemplate(template, FlatToMap(args...))
}

// FlatToMap 集合转map
func FlatToMap(args ...any) map[string]any {
	argMap := make(map[string]any, len(args)/2)
	for i := 0; i < len(args)-1; i += 2 {
		if key, ok := args[i].(string); ok {
			argMap[key] = args[i+1]
		}
	}
	return argMap
}
