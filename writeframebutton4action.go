package frameproto

import (
	"io"
)

// WriteFrameButton4Action will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:4:action" name-value pair.
//
// For example, this call:
//
//	var buttonAction string = "post"
//	
//	frameproto.WriteFrameButton4Action(writer, aspectRatio)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:button:4:action" content="post" />
//
// Note that this package provides some constants to use with WriteFrameButton4Action.
// Namely: ButtonActionLink (for "link"), ButtonActionMint (for "mint"), ButtonActionPost (for "post"), and ButtonActionPostRedirect (for "post_redirect").
//
// Which in code would be used as:
//
//	frameproto.WriteFrameButton4Action(writer, frameproto.ButtonActionLink)
//
// And:
//
//	frameproto.WriteFrameButton4Action(writer, frameproto.ButtonActionMint)
//
// And:
//
//	frameproto.WriteFrameButton4Action(writer, frameproto.ButtonActionPost)
//
// And:
//
//	frameproto.WriteFrameButton4Action(writer, frameproto.ButtonActionPostRedirect)
func WriteFrameButton4Action(writer io.Writer, action string) {
	const property string = MetaPropertyFrameButton4Action
	writeMetaPropertyContent(writer, property, action)
}
