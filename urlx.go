package crawl

import (
	"github.com/cnk3x/urlx"
)

type (
	Request      = urlx.Request
	Option       = urlx.Option
	HeaderOption = urlx.HeaderOption
)

var (
	New            = urlx.New
	AcceptHTML     = urlx.AcceptHTML
	AcceptChinese  = urlx.AcceptChinese
	NoCache        = urlx.NoCache
	UseClient      = urlx.UseClient
	Accept         = urlx.Accept
	AcceptAny      = urlx.AcceptAny
	AcceptJSON     = urlx.AcceptJSON
	AcceptLanguage = urlx.AcceptLanguage
	AcceptXML      = urlx.AcceptXML
	HeaderDel      = urlx.HeaderDel
	HeaderSet      = urlx.HeaderSet
)
