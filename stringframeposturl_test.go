package frameproto

import (
	"testing"
)

func TestStringFramePostURL(t *testing.T) {

	tests := []struct{
		URL string
		Expected string
	}{
		{
			URL: "",
			Expected: `<meta property="fc:frame:post_url" content="" />`+"\n",
		},



		{
			URL: "something",
			Expected: `<meta property="fc:frame:post_url" content="something" />`+"\n",
		},



		{
			URL: "Hello world! 🙂",
			Expected: `<meta property="fc:frame:post_url" content="Hello world! 🙂" />`+"\n",
		},



		{
			URL: "https://example.com/path/to/post/to.php",
			Expected: `<meta property="fc:frame:post_url" content="https://example.com/path/to/post/to.php" />`+"\n",
		},



		{
			URL: "x-proto:apple/banana/cherry",
			Expected: `<meta property="fc:frame:post_url" content="x-proto:apple/banana/cherry" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		actual := StringFramePostURL(test.URL)

		expected := test.Expected

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
