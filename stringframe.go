package frameproto

// StringFrame will return the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame" name-value pair.
//
// For example, this call:
//
//	var version string = "vNext"
//	
//	s = frameproto.StringFrame(version)
//
// Would return this HTML <meta/> element:
//
//	<meta property="fc:frame" content="vNext" />
//
// Note that this package provides some constants to use with StringFrame.
// Namely: VersionVNext (for "vNext").
//
// Which in code would be used as:
//
//	p = frameproto.StringFrame(p, frameproto.VersionVNext)
func StringFrame(version string) string {
	const property string = MetaPropertyFrame
	return stringMetaPropertyContent(property, version)
}
