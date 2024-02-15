package frameproto

// StringFrameImage will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:image" name-value pair.
//
// For example, this call:
//
//	var url string = "https://example.com/images/screen.png"
//	
//	str := frameproto.StringFrameImage(url)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:image" content="https://example.com/images/screen.png" />
func StringFrameImage(url string) string {
	const property string = MetaPropertyFrameImage
	return stringMetaPropertyContent(property, url)
}

