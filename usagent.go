package crawl

import "github.com/cnk3x/urlx"

var (
	MacChromeAgent  = urlx.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
	MacFirefoxAgent = urlx.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:65.0) Gecko/20100101 Firefox/65.0")
	MacSafariAgent  = urlx.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.3 Safari/605.1.15")

	WindowsChromeAgent = urlx.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
	WindowsEdgeAgent   = urlx.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763")
	WindowsIEAgent     = urlx.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko")

	iOSChromeAgent = urlx.UserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_4 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) CriOS/31.0.1650.18 Mobile/11B554a Safari/8536.25")
	iOSSafariAgent = urlx.UserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4")

	AndroidChromeAgent = urlx.UserAgent("Mozilla/5.0 (Linux; Android 11; SM-G9910) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.59 Mobile Safari/537.36")
	AndroidWebkitAgent = urlx.UserAgent("Mozilla/5.0 (Linux; Android 11; SM-G9910) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30")
	AndroidEdgeAgent   = urlx.UserAgent("Mozilla/5.0 (Linux; Android 11; SM-G9910) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Mobile Safari/537.36 Edge/95.0.1020.55")
)

// type BrowserBuilder struct {
// 	Mozilla     string
// 	Platform    []string
// 	Devices     string
// 	AppleWebKit string
// 	Apps        []BrowserApp
// }
//
// type BrowserApp struct {
// 	Name    string
// 	Version string
// }
//
// func (b BrowserBuilder) String() string {
// 	var ua bytes.Buffer
// 	if b.Mozilla == "" {
// 		b.Mozilla = "5.0"
// 	}
// 	ua.WriteString("Mozilla/")
// 	ua.WriteString(b.Mozilla)
//
// 	if len(b.Platform) > 0 || b.Devices != "" {
// 		ua.WriteByte('(')
// 		for i, p := range b.Platform {
// 			if i > 0 {
// 				ua.WriteString("; ")
// 			}
// 			ua.WriteString(p)
// 		}
// 		if b.Devices != "" {
// 			if len(b.Platform) > 0 {
// 				ua.WriteString("; ")
// 			}
// 			ua.WriteString(b.Devices)
// 		}
// 		ua.WriteByte(')')
// 	}
//
// 	if b.AppleWebKit != "" {
// 		ua.WriteString("AppleWebKit/")
// 		ua.WriteString(b.AppleWebKit)
// 		ua.WriteString("(KHTML, like Gecko)")
// 	}
//
// 	for _, app := range b.Apps {
// 		ua.WriteByte(' ')
// 		ua.WriteString(app.Name)
// 		ua.WriteByte('/')
// 		ua.WriteString(app.Version)
// 	}
//
// 	return ua.String()
// }
