package crawl

import (
	"context"
	"time"

	"github.com/cnk3x/urlx"
)

func Default(ctx context.Context) *urlx.Request {
	return MacEdge(ctx)
}

// NewBrowser 浏览器
func NewBrowser(ctx context.Context) *urlx.Request {
	ms := time.Millisecond
	return urlx.New(ctx).HeaderWith(urlx.AcceptHTML, urlx.AcceptChinese, urlx.NoCache, AcceptAllEncodings).
		ProcessWith(AutoCharset, Decompression).
		TryAt(ms*300, ms*800, ms*1500)
}

// MacEdge Mac Edge 浏览器
func MacEdge(ctx context.Context) *urlx.Request {
	return NewBrowser(ctx).HeaderWith(urlx.MacEdgeAgent)
}

// WindowsEdge Windows Edge 浏览器
func WindowsEdge(ctx context.Context) *urlx.Request {
	return NewBrowser(ctx).HeaderWith(urlx.WindowsEdgeAgent)
}

// IPhoneEdge IPhone Edge 浏览器
func IPhoneEdge(ctx context.Context) *urlx.Request {
	return NewBrowser(ctx).HeaderWith(urlx.IPhoneAgent)
}

// AndroidEdge Android Edge 浏览器
func AndroidEdge(ctx context.Context) *urlx.Request {
	return NewBrowser(ctx).HeaderWith(AndroidEdgeAgent)
}
