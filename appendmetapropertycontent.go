package frameproto

import (
	"sourcecode.social/reiver/go-htmlescape"
)

const metaBegin string =
	`<meta property="`

const metaMiddle string =
	`" content="`

const metaEnd string =
	`" />` + "\n"

func appendMetaPropertyContent(p []byte, property string, content string) []byte {

	p = append(p, metaBegin...)
	p = append(p, htmlescape.String(property)...)
	p = append(p, metaMiddle...)
	p = append(p, htmlescape.String(content)...)
	p = append(p, metaEnd...)

	return p
}
