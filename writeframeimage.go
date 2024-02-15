package frameproto

import (
	"io"
)

// WriteFrameImage will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:image" name-value pair.
//
// For example, this call:
//
//	var url string = "https://example.com/images/screen.png"
//	
//	frameproto.WriteFrameImage(writer, url)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:image" content="https://example.com/images/screen.png" />
func WriteFrameImage(writer io.Writer, url string) error {
	const property string = MetaPropertyFrameImage
	return writeMetaPropertyContent(writer, property, url)
}

