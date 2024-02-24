package frameproto

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-opt"
)

func TestActionBody(t *testing.T) {

	tests := []struct{
		JSON string
		Expected ActionBody
	}{
		{
			JSON:
`
{
  "untrustedData": {
    "fid": 2,
    "url": "https://fcpolls.com/polls/1",
    "messageHash": "0xd2b1ddc6c88e865a33cb1a565e0058d757042974",
    "timestamp": 1706243218,
    "network": 1,
    "buttonIndex": 2,
    "inputText": "hello world",
    "castId": {
      "fid": 226,
      "hash": "0xa48dd46161d8e57725f5e26e34ec19c13ff7f3b9"
    }
  },
  "trustedData": {
    "messageBytes": "d2b1ddc6c88e865a33cb1a565e0058d757042974..."
  }
}
`,
			Expected: ActionBody{
				ActionBodyUntrustedData{
					FID: opt.Something(uint64(2)),
					URL: opt.Something("https://fcpolls.com/polls/1"),
					MessageHash: opt.Something("0xd2b1ddc6c88e865a33cb1a565e0058d757042974"),
					Timestamp: opt.Something(int64(1706243218)),
					Network: opt.Something(uint64(1)),
					ButtonIndex: opt.Something(uint64(2)),
					InputText: opt.Something("hello world"),
					CastID: ActionBodyCastID{
						FID: opt.Something(uint64(226)),
						Hash: opt.Something("0xa48dd46161d8e57725f5e26e34ec19c13ff7f3b9"),
					},
				},
				ActionBodyTrustedData{
					MessageBytes: opt.Something("d2b1ddc6c88e865a33cb1a565e0058d757042974..."),
				},
			},
		},
	}

	for testNumber, test := range tests {

		var p []byte = []byte(test.JSON)

		var actual ActionBody

		err := json.Unmarshal(p, &actual)

		if nil != err {
			t.Errorf("For test #%d, did not expect to get an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what wa expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}
