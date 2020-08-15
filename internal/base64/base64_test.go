package base64

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Base64EncodeString(t *testing.T) {
	out := Base64Encode("test1234")

	require.Equal(t, "dGVzdDEyMzQ=", out)
}

func Test_Base64DecodeString(t *testing.T) {
	out := Base64Decode("dGVzdDEyMzQ=")

	require.Equal(t, "test1234", out)
}
