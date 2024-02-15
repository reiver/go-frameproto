package frameproto

import (
	"testing"
)

func TestAppendFrameButton4Action(t *testing.T) {

	tests := []struct{
		Label string
		Expected string
	}{
		{
			Label: "",
			Expected: `<meta property="fc:frame:button:4:action" content="" />`+"\n",
		},



		{
			Label: "something",
			Expected: `<meta property="fc:frame:button:4:action" content="something" />`+"\n",
		},



		{
			Label: "Hello world! 🙂",
			Expected: `<meta property="fc:frame:button:4:action" content="Hello world! 🙂" />`+"\n",
		},



		{
			Label: "link",
			Expected: `<meta property="fc:frame:button:4:action" content="link" />`+"\n",
		},
		{
			Label: "mint",
			Expected: `<meta property="fc:frame:button:4:action" content="mint" />`+"\n",
		},
		{
			Label: "post",
			Expected: `<meta property="fc:frame:button:4:action" content="post" />`+"\n",
		},
		{
			Label: "post_redirect",
			Expected: `<meta property="fc:frame:button:4:action" content="post_redirect" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		{
			var buffer [256]byte
                        var p []byte = buffer[0:0]

			p = AppendFrameButton4Action(p, test.Label)

			expected := test.Expected
			actual := string(p)

			if expected != actual {
				t.Errorf("For test #%d, the actual written meta-tag is not what was expected." ,testNumber)
				t.Logf("EXPECTED: %s", expected)
				t.Logf("ACTUAL:   %s", actual)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("LABEL: %q", test.Label)
				continue
			}
		}

		{
			const top    string = "<html>\n<head>\n"
			const bottom string = "</head>\n<body>\n</body>\n</html>\n"

			var buffer [256]byte
                        var p []byte = buffer[0:0]

			p = append(p, top...)

			p = AppendFrameButton4Action(p, test.Label)

			p = append(p, bottom...)

			expected := top + test.Expected + bottom
			actual := string(p)

			if expected != actual {
				t.Errorf("For test #%d, the actual written meta-tag is not what was expected." ,testNumber)
				t.Logf("EXPECTED: %s", expected)
				t.Logf("ACTUAL:   %s", actual)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("LABEL: %q", test.Label)
				continue
			}
		}
	}
}
