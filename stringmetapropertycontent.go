package frameproto

func stringMetaPropertyContent(property string, content string) string {

	var buffer [bufferSize]byte
	var p []byte = buffer[0:0]

	p = appendMetaPropertyContent(p, property, content)

	return string(p)
}
