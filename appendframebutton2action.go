package frameproto

// AppendFrameButton2Action will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:2:action" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ..
//	
//	var buttonAction string = "post"
//	
//	p = frameproto.AppendFrameButton2Action(p, aspectRatio)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:button:2:action" content="post" />
//
// Note that this package provides some constants to use with AppendFrameButton2Action.
// Namely: ButtonActionLink (for "link"), ButtonActionMint (for "mint"), ButtonActionPost (for "post"), and ButtonActionPostRedirect (for "post_redirect").
//
// Which in code would be used as:
//
//	p = frameproto.AppendFrameButton2Action(p, frameproto.ButtonActionLink)
//
// And:
//
//	p = frameproto.AppendFrameButton2Action(p, frameproto.ButtonActionMint)
//
// And:
//
//	str :+ frameproto.AppendFrameButton2Action(p, frameproto.ButtonActionPost)
//
// And:
//
//	p = frameproto.AppendFrameButton2Action(p, frameproto.ButtonActionPostRedirect)
func AppendFrameButton2Action(p []byte, action string) []byte {
	const property string = MetaPropertyFrameButton2Action
	return appendMetaPropertyContent(p, property, action)
}
