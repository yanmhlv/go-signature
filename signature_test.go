package signature

import (
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	validTestCases = []struct {
		input    string
		expected string
	}{
		{`test`, `V3bEXE72v0fbrKNV6bHW945WgsuRvkKcUBA4+C2AiUp+ssrItfk4btBhQfuupBsUXY11JZY+2DSK15NlXeziZbtGQRbgq6v/Ou2tWd652r39nKCfwpxLw1foXaLiSP5ZPudi42M7ANG4XygjaErPiIMHHQaEA6WKqSLIPmJmJJ1u6PkLulpG5CEMO+/6tWRBMMd5vZuHxqc2koaCWI6a3sW+75INSPXbSAEnwVkWPIeUg+EYCdUE7RARSlT/xf5Bp37qtEGZSG7TF7198HDQjAMpC7HBY+tUXsNZwXJJpcRxRh0Pqfz7Q9LG0ikFfdQE+0gCnuLOmx0IMtLIFjVQpA==`},
		{`example`, `KzpO4V0hQzF5QXl6VxBurWImpi1XhY0/nfezuaNJd3IE1YvIX1uRqgVcKRRLOcBVDg9uT7oQNOyuBSIYmNqc6YR3IJyWNPol0wPObBi71fpDtDxIuZmAW6CHsoKgjSlP3NPP7X4CheJW/xaVgWW2hveAcV4mXvbY2cIVywVeCE6y+LrKI+doq1sRXAx7akQsAe5aQvFNo419zme682eddwkgESDdlJY2vCQ7mFxpxrnXQ3kqefq6JuGizGwgGLaqsp79hu9rc6Bnkvevhq0LcyLZ4SnHfVnv6sxZGyPEFQLslgk/jtww+IXXna6pI0ki2nxSc4UVbFTUIBL4B3lm/A==`},
		{`{"test":"test"}`, `LMYX772LMBO+r83MakogTLoYnUmkR5TTSCMbzbzWPQBgqakoUwjsQnGAbIvA2ZEZXKgbygEW32crr/OzkJbXJp/XSeXPg93IvlOOiaItcPZAIx2MvBh7tY2vtcNMfodlrEvvG5ySDJQWO4LD/Gv7v6dpJJEsVy7AFdT8Azix2MUGcsljIJFWclYsHjJ96OCm5z4RZeA4tp49QCMKb3AL0TZojxxPL/9vSO//o4IcGQCQ02Sw7/p1c8EogCKI9AcMA2fmcog5fy853wMoZJem+qzLtQfBNKLmmKvFE02V7AkRdPn4q7ilCFzkpmN3T5Rhmazv5JHUIvjokR9c0p/4UQ==`},
		{`{"game_id":132,"request_uuid":"583c985f-fee6-4c0e-bbf5-308aad6265af","token":"55b7518e-b89e-11e7-81be-58404eea6d16"}`, `Kd+/B1NealiUiv/9cI/0MQuwnLiAMteKFesys/b8Koe9pVP/H7Hw54W99+q1uMGizaXj+nMIzcwerFSlSkMj94uqXHueGFvDKI4YGKqntlj7EvID1B7P+VlS/A5RN4RjghIMR3MGnsJZT43G8tAju+xJCzjzDmgS25IPVIZobabIpct87ReqxYfkqIlqgH/uKkpU0ezG25mmhMa82Umat1eu88dJDCa1NsbX9SF5gtdC+A8pYS/o87s2RWHG5VVYM8awAwxPnwZacyKIEbXS59BcAI6StUm+/sJWvSKKvR6lxCiJyQOWzG1IwN9NxBthp5AfQx23G5aDMnYDkSUsIw==`},
		{`{"user":"3nYTOSjdlF6UTz9Ir","country":"XX","currency":"BTC","operator_id":1,"token":"cd6bd8560f3bb8f84325152101adeb45","platform":"GPL_DESKTOP","game_id":39,"lang":"en","lobby_url":"https://examplecasino.io","ip":"::ffff:10.0.0.39"}`, `bL7uNP1K3S0HG8IOC0A5Gf/Cl+Hs3YCVfA0ZrjPgGJFnOstxshCQHB7JbeBhTEDhsqd6CFj4U5xOjzselFkO1HhFrTWssB7CNiXaNmizYp2NKuZhkJcrTswVlk8z9NzAkYJfcqnXiC6lMX1X5t6/+dOX6rvLlHM7yfo9LzhVjKo1on9JMHoW8AiYcC8clKEqpyWTQ70euGXnqxRay5RVAmD1sxOlmz8VIX5irtpMOugNDIL1G3g4IgauPk8T2IfVierOFeALQrNx88Es6Dl8Bgb9ogm1W4xgL3Ve01p59DQNt0oorm0LZt/YqkWYGLL2lpd5Qb1FiX4O7+hfyPKN1Q==`},
	}

	invalidTestCases = []struct {
		input    string
		expected string
	}{
		{`invalid_test`, `V3bEXE72v0fbrKNV6bHW945WgsuRvkKcUBA4+C2AiUp+ssrItfk4btBhQfuupBsUXY11JZY+2DSK15NlXeziZbtGQRbgq6v/Ou2tWd652r39nKCfwpxLw1foXaLiSP5ZPudi42M7ANG4XygjaErPiIMHHQaEA6WKqSLIPmJmJJ1u6PkLulpG5CEMO+/6tWRBMMd5vZuHxqc2koaCWI6a3sW+75INSPXbSAEnwVkWPIeUg+EYCdUE7RARSlT/xf5Bp37qtEGZSG7TF7198HDQjAMpC7HBY+tUXsNZwXJJpcRxRh0Pqfz7Q9LG0ikFfdQE+0gCnuLOmx0IMtLIFjVQpA==`},
		{`invalid_example`, `KzpO4V0hQzF5QXl6VxBurWImpi1XhY0/nfezuaNJd3IE1YvIX1uRqgVcKRRLOcBVDg9uT7oQNOyuBSIYmNqc6YR3IJyWNPol0wPObBi71fpDtDxIuZmAW6CHsoKgjSlP3NPP7X4CheJW/xaVgWW2hveAcV4mXvbY2cIVywVeCE6y+LrKI+doq1sRXAx7akQsAe5aQvFNo419zme682eddwkgESDdlJY2vCQ7mFxpxrnXQ3kqefq6JuGizGwgGLaqsp79hu9rc6Bnkvevhq0LcyLZ4SnHfVnv6sxZGyPEFQLslgk/jtww+IXXna6pI0ki2nxSc4UVbFTUIBL4B3lm/A==`},
		{`{"invalid_test":"test"}`, `LMYX772LMBO+r83MakogTLoYnUmkR5TTSCMbzbzWPQBgqakoUwjsQnGAbIvA2ZEZXKgbygEW32crr/OzkJbXJp/XSeXPg93IvlOOiaItcPZAIx2MvBh7tY2vtcNMfodlrEvvG5ySDJQWO4LD/Gv7v6dpJJEsVy7AFdT8Azix2MUGcsljIJFWclYsHjJ96OCm5z4RZeA4tp49QCMKb3AL0TZojxxPL/9vSO//o4IcGQCQ02Sw7/p1c8EogCKI9AcMA2fmcog5fy853wMoZJem+qzLtQfBNKLmmKvFE02V7AkRdPn4q7ilCFzkpmN3T5Rhmazv5JHUIvjokR9c0p/4UQ==`},
		{`{"invalid_game_id":132,"request_uuid":"583c985f-fee6-4c0e-bbf5-308aad6265af","token":"55b7518e-b89e-11e7-81be-58404eea6d16"}`, `Kd+/B1NealiUiv/9cI/0MQuwnLiAMteKFesys/b8Koe9pVP/H7Hw54W99+q1uMGizaXj+nMIzcwerFSlSkMj94uqXHueGFvDKI4YGKqntlj7EvID1B7P+VlS/A5RN4RjghIMR3MGnsJZT43G8tAju+xJCzjzDmgS25IPVIZobabIpct87ReqxYfkqIlqgH/uKkpU0ezG25mmhMa82Umat1eu88dJDCa1NsbX9SF5gtdC+A8pYS/o87s2RWHG5VVYM8awAwxPnwZacyKIEbXS59BcAI6StUm+/sJWvSKKvR6lxCiJyQOWzG1IwN9NxBthp5AfQx23G5aDMnYDkSUsIw==`},
		{`{"invalid_user":"3nYTOSjdlF6UTz9Ir","country":"XX","currency":"BTC","operator_id":1,"token":"cd6bd8560f3bb8f84325152101adeb45","platform":"GPL_DESKTOP","game_id":39,"lang":"en","lobby_url":"https://examplecasino.io","ip":"::ffff:10.0.0.39"}`, `bL7uNP1K3S0HG8IOC0A5Gf/Cl+Hs3YCVfA0ZrjPgGJFnOstxshCQHB7JbeBhTEDhsqd6CFj4U5xOjzselFkO1HhFrTWssB7CNiXaNmizYp2NKuZhkJcrTswVlk8z9NzAkYJfcqnXiC6lMX1X5t6/+dOX6rvLlHM7yfo9LzhVjKo1on9JMHoW8AiYcC8clKEqpyWTQ70euGXnqxRay5RVAmD1sxOlmz8VIX5irtpMOugNDIL1G3g4IgauPk8T2IfVierOFeALQrNx88Es6Dl8Bgb9ogm1W4xgL3Ve01p59DQNt0oorm0LZt/YqkWYGLL2lpd5Qb1FiX4O7+hfyPKN1Q==`},
	}
)

func TestSign(t *testing.T) {
	assert.FileExists(t, "testdata/private.pem")
	assert.FileExists(t, "testdata/public.pem")

	privKeyData, err := ioutil.ReadFile("testdata/private.pem")
	assert.Nil(t, err)

	pubKeyData, err := ioutil.ReadFile("testdata/public.pem")
	assert.Nil(t, err)

	privKey, err := ParsePrivateKey(privKeyData)
	assert.Nil(t, err)

	pubKey, err := ParsePublicKey(pubKeyData)
	assert.Nil(t, err)

	for _, testCase := range validTestCases {
		testCase := testCase
		t.Run("", func(t *testing.T) {
			sign, err := Sign(privKey, []byte(testCase.input))
			assert.Nil(t, err)

			encodedSign := base64.StdEncoding.EncodeToString(sign)

			assert.Equal(t, testCase.expected, encodedSign)
			assert.Nil(t, Verify(pubKey, []byte(testCase.input), sign))
		})
	}

	for _, testCase := range invalidTestCases {
		testCase := testCase
		t.Run("", func(t *testing.T) {
			currentSign, err := Sign(privKey, []byte(testCase.input))
			assert.Nil(t, err)

			encodedcurrentSign := base64.StdEncoding.EncodeToString(currentSign)

			assert.NotEqual(t, testCase.expected, encodedcurrentSign)

			invalidSign, err := base64.StdEncoding.DecodeString(testCase.expected)
			assert.Nil(t, err)

			assert.NotEqual(t, currentSign, invalidSign)
			assert.NotNil(t, Verify(pubKey, []byte(testCase.input), invalidSign))
		})
	}
}
