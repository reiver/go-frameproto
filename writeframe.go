package frameproto

import (
	"io"
)

// WriteFrame will write the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame" name-value pair.
//
// For example, this call:
//
//	var version string = "vNext"
//	
//	str := frameproto.WriteFrame(writer, version)
//
// Would write this HTML <meta/> element:
//
//	<meta property="fc:frame" content="vNext" />
//
// Note that this package provides some constants to use with WriteFrame.
// Namely: VersionVNext (for "vNext").
//
// Which in code would be used as:
//
//	str := frameproto.WriteFrame(writer, frameproto.VersionVNext)
func WriteFrame(writer io.Writer, version string) error {
	const property string = MetaPropertyFrame
	return writeMetaPropertyContent(writer, property, version)
}
