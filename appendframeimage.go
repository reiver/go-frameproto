package frameproto

// AppendFrameImage will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:image" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ...
//	
//	var url string = "https://example.com/images/screen.png"
//	
//	p = frameproto.AppendFrameImage(p, url)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:image" content="https://example.com/images/screen.png" />
func AppendFrameImage(p []byte, url string) []byte {
	const property string = MetaPropertyFrameImage
	return appendMetaPropertyContent(p, property, url)
}
