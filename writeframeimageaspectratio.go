package frameproto

import (
	"io"
)

// WriteFrameImageAspectRatio will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:image:aspect_ratio" name-value pair.
//
// For example, this call:
//
//	var aspectRatio string = "1.91:1"
//	
//	frameproto.WriteFrameImageAspectRatio(writer, aspectRatio)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame:image:aspect_ratio" content="1.91:1" />
//
// Note that this package provides some constants to use with WriteFrameImageAspectRatio.
// Namely: AspectRatioOnePointNineOneToOne (for "1.91:1") and AspectRatioOneToOne (for "1:1").
//
// Which in code would be used at:
//
//	frameproto.WriteFrameImageAspectRatio(writer, frameproto.AspectRatioOnePointNineOneToOne)
//
// And:
//
//	frameproto.WriteFrameImageAspectRatio(writer, frameproto.AspectRatioOneToOne)
func WriteFrameImageAspectRatio(writer io.Writer, label string) error {
	const property string = MetaPropertyFrameImageAspectRatio
	return writeMetaPropertyContent(writer, property, label)
}
