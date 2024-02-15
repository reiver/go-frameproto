package frameproto

// StringFrameButton4Action will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:4:action" name-value pair.
//
// For example, this call:
//
//	var buttonAction string = "post"
//	
//	str := frameproto.StringFrameButton4Action(aspectRatio)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:button:4:action" content="post" />
//
// Note that this package provides some constants to use with StringFrameButton4Action.
// Namely: ButtonActionLink (for "link"), ButtonActionMint (for "mint"), ButtonActionPost (for "post"), and ButtonActionPostRedirect (for "post_redirect").
//
// Which in code would be used as:
//
//	str := frameproto.StringFrameButton4Action(frameproto.ButtonActionLink)
//
// And:
//
//	str := frameproto.StringFrameButton4Action(frameproto.ButtonActionMint)
//
// And:
//
//	str :+ frameproto.StringFrameButton4Action(frameproto.ButtonActionPost)
//
// And:
//
//	str := frameproto.StringFrameButton4Action(frameproto.ButtonActionPostRedirect)
func StringFrameButton4Action(action string) string {
	const property string = MetaPropertyFrameButton4Action
	return stringMetaPropertyContent(property, action)
}
