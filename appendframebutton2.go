package frameproto

// AppendFrameButton2 will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:2" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ...
//	
//	var label string = "go forward"
//	
//	p = frameproto.AppendFrameButton2(p, label)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:button:2" content="go forward" />
func AppendFrameButton2(p []byte, label string) []byte {
	const property string = MetaPropertyFrameButton2
	return appendMetaPropertyContent(p, property, label)
}
