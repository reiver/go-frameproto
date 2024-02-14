package frameproto

import (
	"testing"
)

func TestAppendFramePostURL(t *testing.T) {

	tests := []struct{
		Version string
		Expected string
	}{
		{
			Version:                                              "",
			Expected: `<meta property="fc:frame:post_url" content="" />`+"\n",
		},



		{
			Version:                                              "something",
			Expected: `<meta property="fc:frame:post_url" content="something" />`+"\n",
		},



		{
			Version:                                              "Hello world! 🙂",
			Expected: `<meta property="fc:frame:post_url" content="Hello world! 🙂" />`+"\n",
		},



		{
			Version:                                              "https://example.com/path/to/post/to.php",
			Expected: `<meta property="fc:frame:post_url" content="https://example.com/path/to/post/to.php" />`+"\n",
		},



		{
			Version:                                              "x-proto:apple/banana/cherry",
			Expected: `<meta property="fc:frame:post_url" content="x-proto:apple/banana/cherry" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		{
			var buffer [256]byte
			var p []byte = buffer[0:0]

			p = AppendFramePostURL(p, test.Version)

			expected := test.Expected
			actual   := string(p)

			if expected != actual {
				t.Errorf("For test #%d, the actual written meta-tag is not what was expected.", testNumber)
				t.Logf("EXPECTED: %s", expected)
				t.Logf("ACTUAL:   %s", actual)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("LABEL: %q", test.Version)
				continue
			}
		}

		{
			const top    string = "<html>\n<head>\n"
			const bottom string = "</head>\n<body>\n</body>\n</html>\n"

			var buffer [256]byte
			var p []byte = buffer[0:0]

			p = append(p, top...)

			p = AppendFramePostURL(p, test.Version)

			p = append(p, bottom...)

			expected := top + test.Expected + bottom
			actual   := string(p)

			if expected != actual {
				t.Errorf("For test #%d, the actual written meta-tag is not what was expected.", testNumber)
				t.Logf("EXPECTED: %s", expected)
				t.Logf("ACTUAL:   %s", actual)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("LABEL: %q", test.Version)
				continue
			}
		}
	}
}
