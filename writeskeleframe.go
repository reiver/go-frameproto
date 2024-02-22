package frameproto

import (
	"io"
)

const skeleFrameTop =
	`<!DOCTYPE html>`+"\n"+
	`<html xmlns="http://www.w3.org/1999/xhtml">`+"\n"+
	`<head>`+"\n"+
	`<meta charset="utf-8" />`+"\n"

const skeleFrameBotton =
	`</head>`+"\n"+
	`<body>`+"\n"+
	`</body>`+"\n"+
	`</html>`+"\n"

// WriteSkeleFrame takes care of creating a skeleton-HTML-page that the Frame-Protocol's (i.e., Farcaster Frame's) tags are included in.
//
// For example, this:
//
//	frameproto.WriteSkeleFrame(w, func(writer io.Writer){
//		
//		// This will write:
//		// <meta property="fc:frame" content="vNext" />
//		frameproto.WriteFrame(writer, frameproto.VersionVNext)
//		
//		// This will write:
//		// <meta property="fc:frame:image" content="https://example.com/path/to/img.png" />
//		var imgURL string = "https://example.com/path/to/img.png"
//		frameproto.WriteFrameImage(writer, imgURL)
//	})
//
// Would write this:
//
//	<!DOCTYPE html>
//	<html xmlns="http://www.w3.org/1999/xhtml">
//	<head>
//	<meta charset="utf-8" /
//	<meta property="fc:frame" content="vNext" />
//	<meta property="fc:frame:image" content="https://example.com/path/to/img.png" />
//	</head>
//	<body>
//	</body>
//	</html>
//
// Note that http.ResponseWriter is an io.Writer. So this function can be used to write to http.ResponseWriter.
func WriteSkeleFrame(writer io.Writer, frameWriterFunc func(io.Writer)error) error {

	var err error

	_, err = io.WriteString(writer, skeleFrameTop)
	if nil != err {
		return err
	}

	err = frameWriterFunc(writer)
	if nil != err {
		return err
	}

	_, err = io.WriteString(writer, skeleFrameBotton)
	if nil != err {
		return err
	}

	return nil
}

