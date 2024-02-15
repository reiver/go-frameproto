package frameproto

// AppendFrameButton1 will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:1" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ...
//	
//	var label string = "go forward"
//	
//	p = frameproto.AppendFrameButton1(p, label)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:button:1" content="go forward" />
func AppendFrameButton1(p []byte, label string) []byte {
	const property string = MetaPropertyFrameButton1
	return appendMetaPropertyContent(p, property, label)
}
