package frameproto

import (
	"io"
)

// WriteFrameButton2 will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:2" name-value pair.
//
// For example, this call:
//
//	var label string = "go forward"
//	
//	frameproto.WriteFrameButton2(writer, label)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:button:2" content="go forward" />
func WriteFrameButton2(writer io.Writer, label string) {
	const property string = MetaPropertyFrameButton2
	writeMetaPropertyContent(writer, property, label)
}
