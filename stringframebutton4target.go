package frameproto

// StringFrameButton4Target will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:4:target" name-value pair.
//
// For example, this call:
//
//	var target string = "https://example.com/thing/do-it"
//	
//	str := frameproto.StringFrameButton4Target(target)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:button:4:target" content="https://example.com/thing/do-it" />
func StringFrameButton4Target(target string) string {
	const property string = MetaPropertyFrameButton4Target
	return stringMetaPropertyContent(property, target)
}
