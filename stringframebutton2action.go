package frameproto

// StringFrameButton2Action will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:2:action" name-value pair.
//
// For example, this call:
//
//	var buttonAction string = "post"
//	
//	str := frameproto.StringFrameButton2Action(buttonAction)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:button:2:action" content="post" />
//
// Note that this package provides some constants to use with StringFrameButton2Action.
// Namely: ButtonActionLink (for "link"), ButtonActionMint (for "mint"), ButtonActionPost (for "post"), and ButtonActionPostRedirect (for "post_redirect").
//
// Which in code would be used as:
//
//	// <meta property="fc:frame:button:2:action" content="link" />
//	str := frameproto.StringFrameButton2Action(frameproto.ButtonActionLink)
//
// And:
//
//	// <meta property="fc:frame:button:2:action" content="mint" />
//	str := frameproto.StringFrameButton2Action(frameproto.ButtonActionMint)
//
// And:
//
//	// <meta property="fc:frame:button:2:action" content="post" />
//	str :+ frameproto.StringFrameButton2Action(frameproto.ButtonActionPost)
//
// And:
//
//	// <meta property="fc:frame:button:2:action" content="post_redirect" />
//	str := frameproto.StringFrameButton2Action(frameproto.ButtonActionPostRedirect)
func StringFrameButton2Action(action string) string {
	const property string = MetaPropertyFrameButton2Action
	return stringMetaPropertyContent(property, action)
}
