package frameproto

// These constants are 'names' in name-value pairs used with HTML <meta /> elements by the frame-protocol.
//
// For example:
//
//	<meta property="fc:frame"       content="vNext" />
//	<meta property="fc:frame:image" content="https://example.com/path/to/image.png" />
//	<meta property="og:image"       content="https://example.com/path/to/image.png" />
const (
	MetaPropertyFrame                 = "fc:frame"
	MetaPropertyFrameButton1          = "fc:frame:button:1"
	MetaPropertyFrameButton1Action    = "fc:frame:button:1:action"
	MetaPropertyFrameButton1Target    = "fc:frame:button:1:target"
	MetaPropertyFrameButton2          = "fc:frame:button:2"
	MetaPropertyFrameButton2Action    = "fc:frame:button:2:action"
	MetaPropertyFrameButton2Target    = "fc:frame:button:2:target"
	MetaPropertyFrameButton3          = "fc:frame:button:3"
	MetaPropertyFrameButton3Action    = "fc:frame:button:3:action"
	MetaPropertyFrameButton3Target    = "fc:frame:button:3:target"
	MetaPropertyFrameButton4          = "fc:frame:button:4"
	MetaPropertyFrameButton4Action    = "fc:frame:button:4:action"
	MetaPropertyFrameButton4Target    = "fc:frame:button:4:target"
	MetaPropertyFrameImage            = "fc:frame:image"
	MetaPropertyFrameImageAspectRatio = "fc:frame:image:aspect_ratio"
	MetaPropertyFrameInputText        = "fc:frame:input:text"
	MetaPropertyFramePostURL          = "fc:frame:post_url"
)
