package frameproto

import (
	"testing"
)

func TestStringFrameImageAspectRatio(t *testing.T) {

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

		actual := StringFrameImageAspectRatio(test.AspectRatio)

		expected := test.Expected

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
