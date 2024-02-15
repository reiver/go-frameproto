package frameproto

// StringFrameButton2 will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:2" name-value pair.
//
// For example, this call:
//
//	var label string = "go forward"
//	
//	str := frameproto.StringFrameButton2(label)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:button:2" content="go forward" />
func StringFrameButton2(label string) string {
	const property string = MetaPropertyFrameButton2
	return stringMetaPropertyContent(property, label)
}
