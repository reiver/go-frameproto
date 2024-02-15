package frameproto

// StringFrameButton3Target will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:3:target" name-value pair.
//
// For example, this call:
//
//	var target string = "https://example.com/thing/do-it"
//	
//	str := frameproto.StringFrameButton3Target(target)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:button:3:target" content="https://example.com/thing/do-it" />
func StringFrameButton3Target(target string) string {
	const property string = MetaPropertyFrameButton3Target
	return stringMetaPropertyContent(property, target)
}
