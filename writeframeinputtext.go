package frameproto

import (
	"io"
)

// WriteFrameInputText will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:input:text" name-value pair.
//
// For example, this call:
//
//	var label string = "enter your username"
//	
//	frameproto.WriteFrameInputText(writer, label)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:input:text" content="enter your username" />
func WriteFrameInputText(writer io.Writer, label string) error {
	const property string = MetaPropertyFrameInputText
	return writeMetaPropertyContent(writer, property, label)
}
