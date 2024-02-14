package frameproto

import (
	"io"
)

// WriteFrameButton1Action will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:1:action" name-value pair.
//
// For example, this call:
//
//	var buttonAction string = "post"
//	
//	frameproto.WriteFrameButton1Action(writer, aspectRatio)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:button:1:action" content="post" />
//
// Note that this package provides some constants to use with WriteFrameButton1Action.
// Namely: ButtonActionLink (for "link"), ButtonActionMint (for "mint"), ButtonActionPost (for "post"), and ButtonActionPostRedirect (for "post_redirect").
//
// Which in code would be used as:
//
//	frameproto.WriteFrameButton1Action(writer, frameproto.ButtonActionLink)
//
// And:
//
//	frameproto.WriteFrameButton1Action(writer, frameproto.ButtonActionMint)
//
// And:
//
//	frameproto.WriteFrameButton1Action(writer, frameproto.ButtonActionPost)
//
// And:
//
//	frameproto.WriteFrameButton1Action(writer, frameproto.ButtonActionPostRedirect)
func WriteFrameButton1Action(writer io.Writer, action string) {
	const property string = MetaPropertyFrameButton1Action
	writeMetaPropertyContent(writer, property, action)
}
