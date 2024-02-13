package frameproto

import (
	"io"
)

// WriteFramePostURL will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:post_url" name-value pair.
//
// For example, this call:
//
//	var url string = "https://example.com/my/post/path.php"
//	
//	WriteFramePostURL(writer, url)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:post_url" content="https://example.com/my/post/path.php" />
func WriteFramePostURL(writer io.Writer, url string) {
	const property string = MetaPropertyFramePostURL
	writeMetaPropertyContent(writer, property, url)
}
