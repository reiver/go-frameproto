package frameproto

import (
	"io"
)

type internalFrameWriter struct {
	writer io.Writer
}

var _ FrameWriter = internalFrameWriter{}

func CreateFrameWriter(writer io.Writer) FrameWriter {
	return internalFrameWriter{
		writer:writer,
	}
}



func (receiver internalFrameWriter) WriteFrame(version string) error {
	return WriteFrame(receiver.writer, version)
}



func (receiver internalFrameWriter) WriteFrameButton1(label string) error {
	return WriteFrameButton1(receiver.writer, label)
}

func (receiver internalFrameWriter) WriteFrameButton1Action(label string) error {
	return WriteFrameButton1Action(receiver.writer, label)
}

func (receiver internalFrameWriter) WriteFrameButton1Target(target string) error {
	return WriteFrameButton1Target(receiver.writer, target)
}



func (receiver internalFrameWriter) WriteFrameButton2(label string) error {
	return WriteFrameButton2(receiver.writer, label)
}

func (receiver internalFrameWriter) WriteFrameButton2Action(label string) error {
	return WriteFrameButton2Action(receiver.writer, label)
}

func (receiver internalFrameWriter) WriteFrameButton2Target(target string) error {
	return WriteFrameButton2Target(receiver.writer, target)
}



func (receiver internalFrameWriter) WriteFrameButton3(label string) error {
	return WriteFrameButton3(receiver.writer, label)
}

func (receiver internalFrameWriter) WriteFrameButton3Action(label string) error {
	return WriteFrameButton3Action(receiver.writer, label)
}

func (receiver internalFrameWriter) WriteFrameButton3Target(target string) error {
	return WriteFrameButton3Target(receiver.writer, target)
}



func (receiver internalFrameWriter) WriteFrameButton4(label string) error {
	return WriteFrameButton4(receiver.writer, label)
}

func (receiver internalFrameWriter) WriteFrameButton4Action(label string) error {
	return WriteFrameButton4Action(receiver.writer, label)
}

func (receiver internalFrameWriter) WriteFrameButton4Target(target string) error {
	return WriteFrameButton4Target(receiver.writer, target)
}



func (receiver internalFrameWriter) WriteFrameImage(url string) error {
	return WriteFrameImage(receiver.writer, url)
}

func (receiver internalFrameWriter) WriteFrameImageAspectRatio(aspectRatio string) error {
	return WriteFrameImageAspectRatio(receiver.writer, aspectRatio)
}



func (receiver internalFrameWriter) WriteFrameInputText(label string) error {
	return WriteFrameInputText(receiver.writer, label)
}



func (receiver internalFrameWriter) WriteFramePostURL(url string) error {
	return WriteFramePostURL(receiver.writer, url)
}
