package frameproto

import (
	"testing"
)

func TestAppendFrameImageAspectRatio(t *testing.T) {

	tests := []struct{
		AspectRatio string
		Expected string
	}{
		{
			AspectRatio:                                                    "",
			Expected: `<meta property="fc:frame:image:aspect_ratio" content="" />`+"\n",
		},



		{
			AspectRatio:                                                    "something",
			Expected: `<meta property="fc:frame:image:aspect_ratio" content="something" />`+"\n",
		},



		{
			AspectRatio:                                                    "Hello world! 🙂",
			Expected: `<meta property="fc:frame:image:aspect_ratio" content="Hello world! 🙂" />`+"\n",
		},



		{
			AspectRatio:                                                    "1.91:1",
			Expected: `<meta property="fc:frame:image:aspect_ratio" content="1.91:1" />`+"\n",
		},



		{
			AspectRatio:                                                    "1:1",
			Expected: `<meta property="fc:frame:image:aspect_ratio" content="1:1" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		{
			var buffer [256]byte
			var p []byte = buffer[0:0]

			p = AppendFrameImageAspectRatio(p, test.AspectRatio)

			expected := test.Expected
			actual := string(p)

			if expected != actual {
				t.Errorf("For test #%d, the actual written meta-tag is not what was expected." ,testNumber)
				t.Logf("EXPECTED: %s", expected)
				t.Logf("ACTUAL:   %s", actual)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("AspectRatio: %q", test.AspectRatio)
				continue
			}
		}

		{
			const top    string = "<html>\n<head>\n"
			const bottom string = "</head>\n<body>\n</body>\n</html>\n"

			var buffer [256]byte
			var p []byte = buffer[0:0]

			p = append(p, top...)

			p = AppendFrameImageAspectRatio(p, test.AspectRatio)

			p = append(p, bottom...)

			expected := top + test.Expected + bottom
			actual := string(p)

			if expected != actual {
				t.Errorf("For test #%d, the actual written meta-tag is not what was expected." ,testNumber)
				t.Logf("EXPECTED: %s", expected)
				t.Logf("ACTUAL:   %s", actual)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("AspectRatio: %q", test.AspectRatio)
				continue
			}
		}
	}
}
