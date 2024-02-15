package frameproto

// AppendFrameButton4 will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:4" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ...
//	
//	var label string = "go forward"
//	
//	p = frameproto.AppendFrameButton4(p, label)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:button:4" content="go forward" />
func AppendFrameButton4(p []byte, label string) []byte {
	const property string = MetaPropertyFrameButton4
	return appendMetaPropertyContent(p, property, label)
}
