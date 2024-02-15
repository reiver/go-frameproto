package frameproto

// AppendFrameButton4Action will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:4:action" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ..
//	
//	var buttonAction string = "post"
//	
//	p = frameproto.AppendFrameButton4Action(p, buttonAction)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:button:4:action" content="post" />
//
// Note that this package provides some constants to use with AppendFrameButton4Action.
// Namely: ButtonActionLink (for "link"), ButtonActionMint (for "mint"), ButtonActionPost (for "post"), and ButtonActionPostRedirect (for "post_redirect").
//
// Which in code would be used as:
//
//	// <meta property="fc:frame:button:4:action" content="link" />
//	p = frameproto.AppendFrameButton4Action(p, frameproto.ButtonActionLink)
//
// And:
//
//	// <meta property="fc:frame:button:4:action" content="mint" />
//	p = frameproto.AppendFrameButton4Action(p, frameproto.ButtonActionMint)
//
// And:
//
//	// <meta property="fc:frame:button:4:action" content="post" />
//	str :+ frameproto.AppendFrameButton4Action(p, frameproto.ButtonActionPost)
//
// And:
//
//	// <meta property="fc:frame:button:4:action" content="post_redirect" />
//	p = frameproto.AppendFrameButton4Action(p, frameproto.ButtonActionPostRedirect)
func AppendFrameButton4Action(p []byte, action string) []byte {
	const property string = MetaPropertyFrameButton4Action
	return appendMetaPropertyContent(p, property, action)
}
