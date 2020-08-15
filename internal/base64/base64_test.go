package base64

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EncodeString(t *testing.T) {
	out := Encode("test1234")

	require.Equal(t, "dGVzdDEyMzQ=", out)
}

func Test_DecodeString(t *testing.T) {
	out := Decode("dGVzdDEyMzQ=")

	require.Equal(t, "test1234", out)
}
