package frameproto

// StringFrameButton2Target will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:2:target" name-value pair.
//
// For example, this call:
//
//	var target string = "https://example.com/thing/do-it"
//	
//	str := frameproto.StringFrameButton2Target(target)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:button:2:target" content="https://example.com/thing/do-it" />
func StringFrameButton2Target(target string) string {
	const property string = MetaPropertyFrameButton2Target
	return stringMetaPropertyContent(property, target)
}
