package frameproto

import (
	"testing"

	"strings"
)

func TestWriteFrame(t *testing.T) {

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

		var buffer strings.Builder

		WriteFrame(&buffer, test.Version)

		expected := test.Expected
		actual   := buffer.String()

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
