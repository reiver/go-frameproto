package frameproto

import (
	"testing"

	"strings"
)

func TestWriteFrameButton4Target(t *testing.T) {

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

		var buffer strings.Builder

		WriteFrameButton4Target(&buffer, test.Target)

		expected := test.Expected
		actual   := buffer.String()

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
