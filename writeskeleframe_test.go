
package frameproto

import (
	"testing"

	"io"
	"strings"
)

func TestWriteSkeleFrame(t *testing.T) {

	tests := []struct{
		Head string
		Expected string
	}{
		{
			Head: "",
			Expected:
				`<!DOCTYPE html>`+"\n"+
				`<html xmlns="http://www.w3.org/1999/xhtml" lang="en">`+"\n"+
				`<head>`+"\n"+
				`<meta charset="utf-8" />`+"\n"+
				`</head>`+"\n"+
				`<body>`+"\n"+
				`</body>`+"\n"+
				`</html>`+"\n",
		},



		{
			Head: `<!-- Hello world! -->`+"\n",
			Expected:
				`<!DOCTYPE html>`+"\n"+
				`<html xmlns="http://www.w3.org/1999/xhtml" lang="en">`+"\n"+
				`<head>`+"\n"+
				`<meta charset="utf-8" />`+"\n"+

				`<!-- Hello world! -->`+"\n"+

				`</head>`+"\n"+
				`<body>`+"\n"+
				`</body>`+"\n"+
				`</html>`+"\n",
		},



		{
			Head: `<meta property="hello" content="world" />`+"\n",
			Expected:
				`<!DOCTYPE html>`+"\n"+
				`<html xmlns="http://www.w3.org/1999/xhtml" lang="en">`+"\n"+
				`<head>`+"\n"+
				`<meta charset="utf-8" />`+"\n"+

				`<meta property="hello" content="world" />`+"\n"+

				`</head>`+"\n"+
				`<body>`+"\n"+
				`</body>`+"\n"+
				`</html>`+"\n",
		},



		{
			Head:
				`<meta property="fc:frame" content="vNext" />`+"\n"+
				`<meta property="fc:frame:image" content="https://example.com/path/to/img.png" />`+"\n"+
				`<meta property="og:image" content="https://example.com/path/to/img.png" />`+"\n",
			Expected:
				`<!DOCTYPE html>`+"\n"+
				`<html xmlns="http://www.w3.org/1999/xhtml" lang="en">`+"\n"+
				`<head>`+"\n"+
				`<meta charset="utf-8" />`+"\n"+

				`<meta property="fc:frame" content="vNext" />`+"\n"+
				`<meta property="fc:frame:image" content="https://example.com/path/to/img.png" />`+"\n"+
				`<meta property="og:image" content="https://example.com/path/to/img.png" />`+"\n"+

				`</head>`+"\n"+
				`<body>`+"\n"+
				`</body>`+"\n"+
				`</html>`+"\n",
		},
	}

	for testNumber, test := range tests {

		var buffer strings.Builder

		WriteSkeleFrame(&buffer, func(writer io.Writer)error {
			io.WriteString(writer, test.Head)
			return nil
		})

		expected := test.Expected
		actual   := buffer.String()

		if expected != actual {
			t.Errorf("For test #%d, ", testNumber)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("HEAD:\n%s", test.Head)
			t.Logf("HEAD: %q", test.Head)
			continue
		}
	}
}
