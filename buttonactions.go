package frameproto

// These constants are 'values' in name-value pairs used with HTML <meta /> elements by the frame-protocol, when the name is of the form "fc:frame:button:*:action".
// I.e., when the name is: "fc:frame:button:1:action", "fc:frame:button:2:action", "fc:frame:button:3:action", or "fc:frame:button:4:action".
//
// For example:
//
//      <meta property="fc:frame:button:1:action" content="post" />
//      <meta property="fc:frame:button:2:action" content="post_redirect" />
//      <meta property="fc:frame:button:3:action" content="mint" />
//      <meta property="fc:frame:button:4:action" content="link" />
//
// The way these would be used are:
//
//	frameproto.WriteFrameButton1Action(writer, frameproto.ButtonActionLink)
//
// And:
//
//	p = frameproto.AppendFrameButton2Action(p, frameproto.ButtonActionMint)
//
// And:
//
//	s = frameproto.FrameButton3Action(p, frameproto.ButtonActionPost)
//
// Etc.
const (
	ButtonActionLink         = "link"
	ButtonActionMint         = "mint"
	ButtonActionPost         = "post"
	ButtonActionPostRedirect = "post_redirect"
)
