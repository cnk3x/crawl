package crawl

import (
	"context"
	"time"
)

func Default(ctx context.Context) *Request {
	return MacEdge(ctx)
}

// NewBrowser 浏览器
func NewBrowser(ctx context.Context) *Request {
	ms := time.Millisecond
	return New(ctx).HeaderWith(AcceptHTML, AcceptChinese, NoCache, AcceptAllEncodings).
		ProcessWith(AutoCharset, Decompression).
		TryAt(ms*300, ms*800, ms*1500)
}

// MacEdge Mac Edge 浏览器
func MacEdge(ctx context.Context) *Request {
	return NewBrowser(ctx).HeaderWith(MacEdgeAgent)
}

// WindowsEdge Windows Edge 浏览器
func WindowsEdge(ctx context.Context) *Request {
	return NewBrowser(ctx).HeaderWith(WindowsEdgeAgent)
}

// AndroidEdge Android Edge 浏览器
func AndroidEdge(ctx context.Context) *Request {
	return NewBrowser(ctx).HeaderWith(AndroidEdgeAgent)
}
