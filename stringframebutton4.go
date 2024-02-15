package frameproto

// StringFrameButton4 will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:4" name-value pair.
//
// For example, this call:
//
//	var label string = "go forward"
//	
//	str := frameproto.StringFrameButton4(label)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:button:4" content="go forward" />
func StringFrameButton4(label string) string {
	const property string = MetaPropertyFrameButton4
	return stringMetaPropertyContent(property, label)
}
