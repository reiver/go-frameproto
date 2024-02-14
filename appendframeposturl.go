package frameproto

// AppendFramePostURL will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:post_url" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ...
//	
//	var url string = "https://example.com/my/post/path.php"
//	
//	p = frameproto.AppendFramePostURL(p, url)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:post_url" content="https://example.com/my/post/path.php" />
func AppendFramePostURL(p []byte, url string) []byte {
	const property string = MetaPropertyFramePostURL
	return appendMetaPropertyContent(p, property, url)
}
