package frameproto

import (
	"io"
)

// WriteFrameButton1 will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:1" name-value pair.
//
// For example, this call:
//
//	var label string = "go forward"
//	
//	frameproto.WriteFrameButton1(writer, label)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:button:1" content="go forward" />
func WriteFrameButton1(writer io.Writer, label string) {
	const property string = MetaPropertyFrameButton1
	writeMetaPropertyContent(writer, property, label)
}
