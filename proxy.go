package crawl

import (
	"github.com/cnk3x/urlx"
	"github.com/Qingluan/merkur"
)

// Proxy 使用代理，支持 ss:// ssr:// vmess:// http:// https:// sock5://
func Proxy(proxy string) urlx.Option {
	return func(r *urlx.Request) error {
		if proxy != "" {
			if client := merkur.NewProxyHttpClient(proxy, 10); client != nil {
				r.UseClient(client)
			}
		}
		return nil
	}
}

// ProxySubscribe 使用订阅来当做代理池
func ProxySubscribe(subscribeUri string) urlx.Option {
	pool := merkur.NewProxyPool(merkur.ParseOrder(subscribeUri)...)
	return func(r *urlx.Request) error {
		if client := pool.GetDialer2().ToHttpClient(10); client != nil {
			r.UseClient(client)
		}
		return nil
	}
}
