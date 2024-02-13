package frameproto

import (
	"testing"

	"strings"
)

func TestWriteFrameImage(t *testing.T) {

	tests := []struct{
		URL string
		Expected string
	}{
		{
			URL: "",
			Expected: `<meta property="fc:frame:image" content="" />`+"\n",
		},



		{
			URL: "something",
			Expected: `<meta property="fc:frame:image" content="something" />`+"\n",
		},



		{
			URL: "Hello world! 🙂",
			Expected: `<meta property="fc:frame:image" content="Hello world! 🙂" />`+"\n",
		},



		{
			URL: "https://example.com/path/to/post/to.png",
			Expected: `<meta property="fc:frame:image" content="https://example.com/path/to/post/to.png" />`+"\n",
		},



		{
			URL: "x-proto:apple/banana/cherry",
			Expected: `<meta property="fc:frame:image" content="x-proto:apple/banana/cherry" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		var buffer strings.Builder

		WriteFrameImage(&buffer, test.URL)

		expected := test.Expected
		actual   := buffer.String()

		if expected != actual {
			t.Errorf("For test #%d, the actual written meta-tag is not what was expected." ,testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("URL: %q", test.URL)
			continue
		}
	}
}
