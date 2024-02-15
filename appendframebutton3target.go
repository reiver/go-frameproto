package frameproto

// AppendFrameButton3Target will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:3:target" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ...
//	
//	var target string = "https://example.com/thing/do-it"
//	
//	p = frameproto.AppendFrameButton3Target(p, target)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:button:3:target" content="https://example.com/thing/do-it" />
func AppendFrameButton3Target(p []byte, target string) []byte {
	const property string = MetaPropertyFrameButton3Target
	return appendMetaPropertyContent(p, property, target)
}
