package frameproto

import (
	"testing"

	"strings"
)

func TestWriteMetaPropertyContent(t *testing.T) {

	tests := []struct{
		Property string
		Content string
		Expected string
	}{
		{
			Property:                 "",
			Content:                             "",
			Expected: `<meta property="" content="" />`+"\n",
		},



		{
			Property:                 "name",
			Content:                                 "value",
			Expected: `<meta property="name" content="value" />`+"\n",
		},



		{
			Property:                 "test",
			Content:                                 "5 > 3",
			Expected: `<meta property="test" content="5 &gt; 3" />`+"\n",
		},
		{
			Property:                 "test",
			Content:                                 "5 < 3",
			Expected: `<meta property="test" content="5 &lt; 3" />`+"\n",
		},



		{
			Property:                 "quotation",
			Content:                                      `she said, "hello! 🙂"`,
			Expected: `<meta property="quotation" content="she said, &#34;hello! 🙂&#34;" />`+"\n",
		},



		{
			Property:                 "apple\tbanana\tcherry",
			Content:                                                        "\x00\x01\x02\x03\x04\x05\x06\x07\x08",
			Expected: `<meta property="apple&#x9;banana&#x9;cherry" content="&#x0;&#x1;&#x2;&#x3;&#x4;&#x5;&#x6;&#x7;&#x8;" />`+"\n",
		},
	}

	for testNumber, test := range tests {

		var buffer strings.Builder

		writeMetaPropertyContent(&buffer, test.Property, test.Content)

		expected := test.Expected
		actual   := buffer.String()

		if expected != actual {
			t.Errorf("For test #%d, the actual rendered <meta/> tag is not what was expected." , testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PROPERTY: %s", test.Property)
			t.Logf("CONTENT:  %s", test.Content)
			continue
		}
	}
}
