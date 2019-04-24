package signature

import (
	"io/ioutil"
	"testing"
)

func BenchmarkSign(b *testing.B) {
	privKeyData, err := ioutil.ReadFile("testdata/private.pem")
	if err != nil {
		b.FailNow()
	}
	pubKeyData, err := ioutil.ReadFile("testdata/public.pem")
	if err != nil {
		b.FailNow()
	}

	privKey, err := ParsePrivateKey(privKeyData)
	if err != nil {
		b.FailNow()
	}
	pubKey, err := ParsePublicKey(pubKeyData)
	if err != nil {
		b.FailNow()
	}

	b.ResetTimer()
	for _, testCase := range validTestCases {
		testCase := testCase
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sign, err := Sign(privKey, []byte(testCase.input))
				if err != nil {
					b.FailNow()
				}
				if err := Verify(pubKey, []byte(testCase.input), sign); err != nil {
					b.FailNow()
				}
			}
		})
	}
}
