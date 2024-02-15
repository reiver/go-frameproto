package frameproto

import (
	"io"
)

func writeMetaPropertyContent(writer io.Writer, property string, content string) error {

	if nil == writer {
		return errNilWriter
	}

	var buffer [bufferSize]byte
	var p []byte = buffer[0:0]

	p = appendMetaPropertyContent(p, property, content)

	_, err := writer.Write(p)
	return err
}
