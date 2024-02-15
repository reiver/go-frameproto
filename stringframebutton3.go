package frameproto

// StringFrameButton3 will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:3" name-value pair.
//
// For example, this call:
//
//	var label string = "go forward"
//	
//	str := frameproto.StringFrameButton3(label)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:button:3" content="go forward" />
func StringFrameButton3(label string) string {
	const property string = MetaPropertyFrameButton3
	return stringMetaPropertyContent(property, label)
}
