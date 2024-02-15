package frameproto

import (
	"testing"
)

func TestAppendFrameButton4Target(t *testing.T) {

	tests := []struct{
		Target string
		Expected string
	}{
		{
			Target: "",
			Expected: `<meta property="fc:frame:button:4:target" content="" />`+"\n",
		},



		{
			Target: "something",
			Expected: `<meta property="fc:frame:button:4:target" content="something" />`+"\n",
		},



		{
			Target: "Hello world! 🙂",
			Expected: `<meta property="fc:frame:button:4:target" content="Hello world! 🙂" />`+"\n",
		},



		{
			Target: "https://example.com/path/to/post/to.php",
			Expected: `<meta property="fc:frame:button:4:target" content="https://example.com/path/to/post/to.php" />`+"\n",
		},



		{
			Target: "x-proto:apple/banana/cherry",
			Expected: `<meta property="fc:frame:button:4:target" content="x-proto:apple/banana/cherry" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		{
			var buffer [256]byte
			var p []byte = buffer[0:0]

			p = AppendFrameButton4Target(p, test.Target)

			expected := test.Expected
			actual := string(p)

			if expected != actual {
				t.Errorf("For test #%d, the actual written meta-tag is not what was expected." ,testNumber)
				t.Logf("EXPECTED: %s", expected)
				t.Logf("ACTUAL:   %s", actual)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("Target: %q", test.Target)
				continue
			}
		}

		{
			const top    string = "<html>\n<head>\n"
			const bottom string = "</head>\n<body>\n</body>\n</html>\n"

			var buffer [256]byte
			var p []byte = buffer[0:0]

			p = append(p, top...)

			p = AppendFrameButton4Target(p, test.Target)

			p = append(p, bottom...)

			expected := top + test.Expected + bottom
			actual := string(p)

			if expected != actual {
				t.Errorf("For test #%d, the actual written meta-tag is not what was expected." ,testNumber)
				t.Logf("EXPECTED: %s", expected)
				t.Logf("ACTUAL:   %s", actual)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("Target: %q", test.Target)
				continue
			}
		}
	}
}
