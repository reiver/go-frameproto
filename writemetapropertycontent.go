package frameproto

import (
	"io"
)

func writeMetaPropertyContent(writer io.Writer, property string, content string) {

	if nil == writer {
		return
	}

	var buffer [bufferSize]byte
	var p []byte = buffer[0:0]

	p = appendMetaPropertyContent(p, property, content)

	writer.Write(p)
}
