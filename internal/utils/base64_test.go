package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_base64_encode(t *testing.T) {
	out := Base64Encode("test1234")

	require.Equal(t, "dGVzdDEyMzQ=", out)
}

func Test_base64_decode(t *testing.T) {
	out := Base64Decode("dGVzdDEyMzQ=")

	require.Equal(t, "test1234", out)
}
