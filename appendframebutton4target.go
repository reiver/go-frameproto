package frameproto

// AppendFrameButton4Target will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:button:4:target" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ...
//	
//	var target string = "https://example.com/thing/do-it"
//	
//	p = frameproto.AppendFrameButton4Target(p, target)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:button:4:target" content="https://example.com/thing/do-it" />
func AppendFrameButton4Target(p []byte, target string) []byte {
	const property string = MetaPropertyFrameButton4Target
	return appendMetaPropertyContent(p, property, target)
}
