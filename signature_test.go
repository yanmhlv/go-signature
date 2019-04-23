package signature

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	assert.FileExists(t, "testdata/private.pem")
	assert.FileExists(t, "testdata/public.pub")

	privKeyData, err := ioutil.ReadFile("testdata/private.pem")
	assert.Nil(t, err)

	pubKeyData, err := ioutil.ReadFile("testdata/public.pub")
	assert.Nil(t, err)

	privKey, err := ParsePrivateKey(privKeyData)
	assert.Nil(t, err)

	pubKey, err := ParsePublicKey(pubKeyData)
	assert.Nil(t, err)

	msg := []byte(`hello world`)

	sign, err := Sign(privKey, msg)
	assert.Nil(t, err)
	assert.NotEqual(t, sign, msg)

	assert.Nil(t, Verify(pubKey, msg, sign))

	assert.NotNil(t, Verify(pubKey, msg, []byte(`invalid sign`)))
}
