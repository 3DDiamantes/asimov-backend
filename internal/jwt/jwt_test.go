package jwt

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GenerateToken(t *testing.T) {
	header := Header{
		Algorithm: "HS256",
		Type:      "JWT",
	}
	payload := Payload{
		Subject: "Cosme Fulanito",
	}

	token := GenerateToken(header, payload)

	require.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDb3NtZSBGdWxhbml0byJ9.dRv0cRZfo90TfgiXCFzMqZudLPYxGTD1M-YeMXAzlJs", token)
}
func Test_generateSignature(t *testing.T) {
	header := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9`
	payload := `eyJzdWIiOiJDb3NtZSBGdWxhbml0byJ9`

	signature := generateSignature(header, payload)

	require.Equal(t, "dRv0cRZfo90TfgiXCFzMqZudLPYxGTD1M-YeMXAzlJs", signature)
}

func Test_VerifyOK(t *testing.T) {
	jwt := strings.Split("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDb3NtZSBGdWxhbml0byJ9.dRv0cRZfo90TfgiXCFzMqZudLPYxGTD1M-YeMXAzlJs", ".")

	out := Verify(jwt[0], jwt[1], jwt[2])

	require.True(t, out)
}

func Test_VerifyFail(t *testing.T) {
	jwt := strings.Split("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDb3NtZSBGdWxhbml0byJ8.dRv0cRZfo90TfgiXCFzMqZudLPYxGTD1M-YeMXAzlJs", ".")

	out := Verify(jwt[0], jwt[1], jwt[2])

	require.False(t, out)
}
