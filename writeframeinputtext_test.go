package frameproto

import (
	"testing"

	"strings"
)

func TestWriteFrameInputText(t *testing.T) {

	tests := []struct{
		URL string
		Expected string
	}{
		{
			URL: "",
			Expected: `<meta property="fc:frame:input:text" content="" />`+"\n",
		},



		{
			URL: "something",
			Expected: `<meta property="fc:frame:input:text" content="something" />`+"\n",
		},



		{
			URL: "Hello world! 🙂",
			Expected: `<meta property="fc:frame:input:text" content="Hello world! 🙂" />`+"\n",
		},



		{
			URL: "enter your username",
			Expected: `<meta property="fc:frame:input:text" content="enter your username" />`+"\n",
		},



		{
			URL: "I like to eat, eat, eat, apples and bananas",
			Expected: `<meta property="fc:frame:input:text" content="I like to eat, eat, eat, apples and bananas" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		var buffer strings.Builder

		WriteFrameInputText(&buffer, test.URL)

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
