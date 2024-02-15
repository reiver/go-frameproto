package frameproto

// AppendFrameImageAspectRatio will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:image:aspect_ratio" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ..
//	
//	var aspectRatio string = "1.91:1"
//	
//	p = frameproto.AppendFrameImageAspectRatio(p, aspectRatio)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:image:aspect_ratio" content="1.91:1" />
//
// Note that this package provides some constants to use with AppendFrameImageAspectRatio.
// Namely: AspectRatioOnePointNineOneToOne (for "1.91:1") and AspectRatioOneToOne (for "1:1").
//
// Which in code would be used at:
//
//	p = frameproto.AppendFrameImageAspectRatio(p, frameproto.AspectRatioOnePointNineOneToOne)
//
// And:
//
//	p = frameproto.AppendFrameImageAspectRatio(p, frameproto.AspectRatioOneToOne)
func AppendFrameImageAspectRatio(p []byte, label string) []byte {
	const property string = MetaPropertyFrameImageAspectRatio
	return appendMetaPropertyContent(p, property, label)
}
