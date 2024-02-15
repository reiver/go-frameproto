package frameproto

import (
	"io"
)

// WriteFrameButton1Target will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:1:target" name-value pair.
//
// For example, this call:
//
//	var target string = "https://example.com/thing/do-it"
//	
//	frameproto.WriteFrameButton1Target(writer, target)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:button:1:target" content="https://example.com/thing/do-it" />
func WriteFrameButton1Target(writer io.Writer, target string) {
	const property string = MetaPropertyFrameButton1Target
	writeMetaPropertyContent(writer, property, target)
}
