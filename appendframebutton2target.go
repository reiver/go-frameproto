package frameproto

// AppendFrameButton2Target will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:2:target" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ...
//	
//	var target string = "https://example.com/thing/do-it"
//	
//	p = frameproto.AppendFrameButton2Target(p, target)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:button:2:target" content="https://example.com/thing/do-it" />
func AppendFrameButton2Target(p []byte, target string) []byte {
	const property string = MetaPropertyFrameButton2Target
	return appendMetaPropertyContent(p, property, target)
}
