package frameproto

// AppendFrame will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ...
//	
//	var version string = "vNext"
//	
//	p = frameproto.AppendFrame(p, version)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame" content="vNext" />
//
// Note that this package provides some constants to use with AppendFrame.
// Namely: VersionVNext (for "vNext").
//
// Which in code would be used as:
//
//	p = frameproto.AppendFrame(p, frameproto.VersionVNext)
func AppendFrame(p []byte, version string) []byte {
	const property string = MetaPropertyFrame
	return appendMetaPropertyContent(p, property, version)
}
