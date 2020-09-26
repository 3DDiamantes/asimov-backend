package jwt

import (
	"asimov-backend/internal/defines"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var prevJwtSecret string

func Test_GenerateToken(t *testing.T) {
	setupTest()
	header := Header{
		Algorithm: "HS256",
		Type:      "JWT",
	}
	payload := Payload{
		Subject: "Cosme Fulanito",
	}

	token := GenerateToken(header, payload)

	require.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDb3NtZSBGdWxhbml0byJ9.LHyGHUaNRnctIkDvGuy2Mr5CAhYwpJ_G8U-F9ccoOJA", token)
	teardownTest()
}
func Test_generateSignature(t *testing.T) {
	setupTest()
	header := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9`
	payload := `eyJzdWIiOiJDb3NtZSBGdWxhbml0byJ9`

	signature := generateSignature(header, payload)

	require.Equal(t, "LHyGHUaNRnctIkDvGuy2Mr5CAhYwpJ_G8U-F9ccoOJA", signature)
	teardownTest()
}

func Test_VerifyOK(t *testing.T) {
	setupTest()
	jwt := strings.Split("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDb3NtZSBGdWxhbml0byJ9.LHyGHUaNRnctIkDvGuy2Mr5CAhYwpJ_G8U-F9ccoOJA", ".")

	out := Verify(jwt[0], jwt[1], jwt[2])

	require.True(t, out)
	teardownTest()
}

func Test_VerifyFail(t *testing.T) {
	setupTest()
	jwt := strings.Split("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDb3NtZSBGdWxhbml0byJ8.LHyGHUaNRnctIkDvGuy2Mr5CAhYwpJ_G8U-F9ccoOJA", ".")

	out := Verify(jwt[0], jwt[1], jwt[2])

	require.False(t, out)
	teardownTest()
}

func setupTest() {
	prevJwtSecret = os.Getenv(defines.EnvJWTSecret)
	os.Setenv(defines.EnvJWTSecret, "dGVzdA")
}
func teardownTest() {
	os.Setenv(defines.EnvJWTSecret, prevJwtSecret)
	prevJwtSecret = ""
}
