package frameproto

import (
	"testing"
)

func TestStringFrameButton3(t *testing.T) {

	tests := []struct{
		Label string
		Expected string
	}{
		{
			Label: "",
			Expected: `<meta property="fc:frame:button:3" content="" />`+"\n",
		},



		{
			Label: "something",
			Expected: `<meta property="fc:frame:button:3" content="something" />`+"\n",
		},



		{
			Label: "Hello world! 🙂",
			Expected: `<meta property="fc:frame:button:3" content="Hello world! 🙂" />`+"\n",
		},



		{
			Label: "go forward",
			Expected: `<meta property="fc:frame:button:3" content="go forward" />`+"\n",
		},



		{
			Label: "I like to eat, eat, eat, apples and bananas",
			Expected: `<meta property="fc:frame:button:3" content="I like to eat, eat, eat, apples and bananas" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		actual := StringFrameButton3(test.Label)

		expected := test.Expected

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
