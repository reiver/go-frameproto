package frameproto

import (
	"testing"
)

func TestStringFrameButton2Target(t *testing.T) {

	tests := []struct{
		Target string
		Expected string
	}{
		{
			Target: "",
			Expected: `<meta property="fc:frame:button:2:target" content="" />`+"\n",
		},



		{
			Target: "something",
			Expected: `<meta property="fc:frame:button:2:target" content="something" />`+"\n",
		},



		{
			Target: "Hello world! 🙂",
			Expected: `<meta property="fc:frame:button:2:target" content="Hello world! 🙂" />`+"\n",
		},



		{
			Target: "https://example.com/path/to/post/to.php",
			Expected: `<meta property="fc:frame:button:2:target" content="https://example.com/path/to/post/to.php" />`+"\n",
		},



		{
			Target: "x-proto:apple/banana/cherry",
			Expected: `<meta property="fc:frame:button:2:target" content="x-proto:apple/banana/cherry" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		actual := StringFrameButton2Target(test.Target)

		expected := test.Expected

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
