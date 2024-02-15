package frameproto

// StringFrameButton1Target will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:1:target" name-value pair.
//
// For example, this call:
//
//	var target string = "https://example.com/thing/do-it"
//	
//	str := frameproto.StringFrameButton1Target(target)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:button:1:target" content="https://example.com/thing/do-it" />
func StringFrameButton1Target(target string) string {
	const property string = MetaPropertyFrameButton1Target
	return stringMetaPropertyContent(property, target)
}
