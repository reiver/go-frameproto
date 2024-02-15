package frameproto

// StringFrameButton1 will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:1" name-value pair.
//
// For example, this call:
//
//	var label string = "go forward"
//	
//	str := frameproto.StringFrameButton1(label)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:button:1" content="go forward" />
func StringFrameButton1(label string) string {
	const property string = MetaPropertyFrameButton1
	return stringMetaPropertyContent(property, label)
}
