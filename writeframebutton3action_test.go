package frameproto

import (
	"testing"

	"strings"
)

func TestWriteFrameButton3Action(t *testing.T) {

	tests := []struct{
		Label string
		Expected string
	}{
		{
			Label: "",
			Expected: `<meta property="fc:frame:button:3:action" content="" />`+"\n",
		},



		{
			Label: "something",
			Expected: `<meta property="fc:frame:button:3:action" content="something" />`+"\n",
		},



		{
			Label: "Hello world! 🙂",
			Expected: `<meta property="fc:frame:button:3:action" content="Hello world! 🙂" />`+"\n",
		},



		{
			Label: "link",
			Expected: `<meta property="fc:frame:button:3:action" content="link" />`+"\n",
		},
		{
			Label: "mint",
			Expected: `<meta property="fc:frame:button:3:action" content="mint" />`+"\n",
		},
		{
			Label: "post",
			Expected: `<meta property="fc:frame:button:3:action" content="post" />`+"\n",
		},
		{
			Label: "post_redirect",
			Expected: `<meta property="fc:frame:button:3:action" content="post_redirect" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		var buffer strings.Builder

		WriteFrameButton3Action(&buffer, test.Label)

		expected := test.Expected
		actual   := buffer.String()

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
