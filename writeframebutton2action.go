package frameproto

import (
	"io"
)

// WriteFrameButton2Action will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:2:action" name-value pair.
//
// For example, this call:
//
//	var buttonAction string = "post"
//	
//	frameproto.WriteFrameButton2Action(writer, aspectRatio)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:button:2:action" content="post" />
//
// Note that this package provides some constants to use with WriteFrameButton2Action.
// Namely: ButtonActionLink (for "link"), ButtonActionMint (for "mint"), ButtonActionPost (for "post"), and ButtonActionPostRedirect (for "post_redirect").
//
// Which in code would be used as:
//
//	frameproto.WriteFrameButton2Action(writer, frameproto.ButtonActionLink)
//
// And:
//
//	frameproto.WriteFrameButton2Action(writer, frameproto.ButtonActionMint)
//
// And:
//
//	frameproto.WriteFrameButton2Action(writer, frameproto.ButtonActionPost)
//
// And:
//
//	frameproto.WriteFrameButton2Action(writer, frameproto.ButtonActionPostRedirect)
func WriteFrameButton2Action(writer io.Writer, action string) {
	const property string = MetaPropertyFrameButton2Action
	writeMetaPropertyContent(writer, property, action)
}
