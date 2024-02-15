package frameproto

// AppendFrameButton3 will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:3" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ...
//	
//	var label string = "go forward"
//	
//	p = frameproto.AppendFrameButton3(p, label)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:button:3" content="go forward" />
func AppendFrameButton3(p []byte, label string) []byte {
	const property string = MetaPropertyFrameButton3
	return appendMetaPropertyContent(p, property, label)
}
