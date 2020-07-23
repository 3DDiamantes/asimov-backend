package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"github.com/luxarts/asimov-backend/internal/service"
)

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	svc := service.NewPingService()
	ctrl := NewPingController(svc)
	ctrl.Ping(c)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "pong", w.Body.String())
}
