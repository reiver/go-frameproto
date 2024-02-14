package frameproto

import (
	"testing"
)

func TestAppendFrame(t *testing.T) {

	tests := []struct{
		Version string
		Expected string
	}{
		{
			Version:                                     "",
			Expected: `<meta property="fc:frame" content="" />`+"\n",
		},



		{
			Version:                                      "something",
			Expected: `<meta property="fc:frame" content="something" />`+"\n",
		},



		{
			Version:                                     "Hello world! 🙂",
			Expected: `<meta property="fc:frame" content="Hello world! 🙂" />`+"\n",
		},



		{
			Version:                                     "vNext",
			Expected: `<meta property="fc:frame" content="vNext" />`+"\n",
		},



		{
			Version:                                     "2020-01-01",
			Expected: `<meta property="fc:frame" content="2020-01-01" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		{
			var buffer [256]byte
			var p []byte = buffer[0:0]

			p = AppendFrame(p, test.Version)

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

			p = AppendFrame(p, test.Version)

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
