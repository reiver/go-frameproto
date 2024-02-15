package frameproto

// AppendFrameInputText will append the HTML <meta/> element for the Frame-Protocol's (i.e., Farcaster Frame's) "fc:frame:input:text" name-value pair.
//
// For example, this call:
//
//	var p []byte
//	
//	// ...
//	
//	var label string = "enter your username
//	
//	p = frameproto.AppendFrameInputText(p, label)
//
// Would append this HTML <meta/> element:
//
//	<meta property="fc:frame:input:text" content="enter your username" />
func AppendFrameInputText(p []byte, url string) []byte {
	const property string = MetaPropertyFrameInputText
	return appendMetaPropertyContent(p, property, url)
}
