package test

import (
	"asimov-backend/internal/defines"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPing_OK(t *testing.T) {
	// Given
	header := map[string]string{
		defines.AuthAuthorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCJ9.ml0QsTms3K9wMygTu41ZhKlTyjmW9zHQtoS8FUsCCjU",
	}

	// When
	resp := executeRequest(http.MethodGet, defines.PathPingGet, header, "")

	// Then
	require.Equal(t, http.StatusOK, resp.Code)
	require.Equal(t, "pong", resp.Body.String())
}
func TestPing_InvalidSignature(t *testing.T) {
	// Given
	header := map[string]string{
		defines.AuthAuthorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCJ9.ml0QsTms3K9wMygTu41ZhKlTyjmW9zHQtoS8FUsCCjo",
	}

	// When
	resp := executeRequest(http.MethodGet, defines.PathPingGet, header, "")

	// Then
	require.Equal(t, http.StatusNotFound, resp.Code)
}
func TestPing_NoAuth(t *testing.T) {
	// Given

	// When
	resp := executeRequest(http.MethodGet, defines.PathPingGet, nil, "")

	// Then
	require.Equal(t, http.StatusBadRequest, resp.Code)
}
