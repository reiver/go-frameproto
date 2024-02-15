package frameproto

// StringFramePostURL will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:post_url" name-value pair.
//
// For example, this call:
//
//	var url string = "https://example.com/my/post/path.php"
//	
//	str := frameproto.StringFramePostURL(url)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:post_url" content="https://example.com/my/post/path.php" />
func StringFramePostURL(url string) string {
	const property string = MetaPropertyFramePostURL
	return stringMetaPropertyContent(property, url)
}
