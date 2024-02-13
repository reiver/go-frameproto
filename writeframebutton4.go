package frameproto

import (
	"io"
)

// WriteFrameButton4 will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:4" name-value pair.
//
// For example, this call:
//
//	var label string = "go forward"
//	
//	frameproto.WriteFrameButton4(writer, label)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:button:4" content="go forward" />
func WriteFrameButton4(writer io.Writer, label string) {
	const property string = MetaPropertyFrameButton4
	writeMetaPropertyContent(writer, property, label)
}
