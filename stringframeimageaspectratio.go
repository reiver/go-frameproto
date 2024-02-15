package frameproto

// StringFrameImageAspectRatio will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:image:aspect_ratio" name-value pair.
//
// For example, this call:
//
//	var aspectRatio string = "1.91:1"
//	
//	str := frameproto.StringFrameImageAspectRatio(aspectRatio)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame:image:aspect_ratio" content="1.91:1" />
//
// Note that this package provides some constants to use with StringFrameImageAspectRatio.
// Namely: AspectRatioOnePointNineOneToOne (for "1.91:1") and AspectRatioOneToOne (for "1:1").
//
// Which in code would be used at:
//
//	str := frameproto.StringFrameImageAspectRatio(frameproto.AspectRatioOnePointNineOneToOne)
//
// And:
//
//	str := frameproto.StringFrameImageAspectRatio(frameproto.AspectRatioOneToOne)
func StringFrameImageAspectRatio(label string) string {
	const property string = MetaPropertyFrameImageAspectRatio
	return stringMetaPropertyContent(property, label)
}
