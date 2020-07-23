package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	svc := NewPingService()
	require.Equal(t, svc.Ping(), "pong")
}
