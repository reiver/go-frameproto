package frameproto

import (
	"io"
)

// WriteFrameButton3 will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:3" name-value pair.
//
// For example, this call:
//
//	var label string = "go forward"
//	
//	frameproto.WriteFrameButton3(writer, label)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:button:3" content="go forward" />
func WriteFrameButton3(writer io.Writer, label string) {
	const property string = MetaPropertyFrameButton3
	writeMetaPropertyContent(writer, property, label)
}
