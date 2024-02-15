package frameproto

// StringFrameButton3Action will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:3:action" name-value pair.
//
// For example, this call:
//
//	var buttonAction string = "post"
//	
//	str := frameproto.StringFrameButton3Action(aspectRatio)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:button:3:action" content="post" />
//
// Note that this package provides some constants to use with StringFrameButton3Action.
// Namely: ButtonActionLink (for "link"), ButtonActionMint (for "mint"), ButtonActionPost (for "post"), and ButtonActionPostRedirect (for "post_redirect").
//
// Which in code would be used as:
//
//	str := frameproto.StringFrameButton3Action(frameproto.ButtonActionLink)
//
// And:
//
//	str := frameproto.StringFrameButton3Action(frameproto.ButtonActionMint)
//
// And:
//
//	str :+ frameproto.StringFrameButton3Action(frameproto.ButtonActionPost)
//
// And:
//
//	str := frameproto.StringFrameButton3Action(frameproto.ButtonActionPostRedirect)
func StringFrameButton3Action(action string) string {
	const property string = MetaPropertyFrameButton3Action
	return stringMetaPropertyContent(property, action)
}
