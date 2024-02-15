package frameproto

//
type FrameWriter interface {
	// <meta property="fc:frame" content="{{version}}" />
	WriteFrame(version string) error



	// <meta property="fc:frame:button:1" content="{{label}}" />
	WriteFrameButton1(label string) error

	// <meta property="fc:frame:button:1:action" content="{{action}}" />
	WriteFrameButton1Action(action string) error

	// <meta property="fc:frame:button:1:target" content="{{target}}" />
	WriteFrameButton1Target(buttonTarget string) error



	// <meta property="fc:frame:button:2" content="{{label}}" />
	WriteFrameButton2(label string) error

	// <meta property="fc:frame:button:2:action" content="{{action}}" />
	WriteFrameButton2Action(action string) error

	// <meta property="fc:frame:button:2:target" content="{{target}}" />
	WriteFrameButton2Target(target string) error



	// <meta property="fc:frame:button:3" content="{{label}}" />
	WriteFrameButton3(label string) error

	// <meta property="fc:frame:button:3:action" content="{{action}}" />
	WriteFrameButton3Action(buttonAction string) error

	// <meta property="fc:frame:button:3:target" content="{{target}}" />
	WriteFrameButton3Target(buttonTarget string) error



	// <meta property="fc:frame:button:4" content="{{label}}" />
	WriteFrameButton4(label string) error

	// <meta property="fc:frame:button:4:action" content="{{action}}" />
	WriteFrameButton4Action(action string) error

	// <meta property="fc:frame:button:4:target" content="{{target}}" />
	WriteFrameButton4Target(target string) error



	// <meta property="fc:frame:image" content="{{url}}" />
	WriteFrameImage(url string) error

	// <meta property="fc:frame:image:aspect_ratio" content="{{aspectRatio}}" />
	WriteFrameImageAspectRatio(aspectRatio string) error



	// <meta property="fc:frame:input:text" content="{{label}}" />
	WriteFrameInputText(label string) error



	// <meta property="fc:frame:post_url" content="{{url}}" />
	WriteFramePostURL(url string) error
}
