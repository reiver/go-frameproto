package frameproto

// StringFrameInputText will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:input:text" name-value pair.
//
// For example, this call:
//
//	var label string = "enter your username"
//	
//	str := frameproto.StringFrameInputText(label)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:input:text" content="enter your username" />
func StringFrameInputText(label string) string {
	const property string = MetaPropertyFrameInputText
	return stringMetaPropertyContent(property, label)
}
