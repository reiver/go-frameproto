package frameproto

import (
	"io"
)

// WriteFrameButton3Target will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:3:target" name-value pair.
//
// For example, this call:
//
//	var target string = "https://example.com/thing/do-it"
//	
//	frameproto.WriteFrameButton3Target(writer, target)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:button:3:target" content="https://example.com/thing/do-it" />
func WriteFrameButton3Target(writer io.Writer, url string) {
	const property string = MetaPropertyFrameButton3Target
	writeMetaPropertyContent(writer, property, url)
}
