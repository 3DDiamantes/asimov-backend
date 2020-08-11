package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type PingServiceMock struct {
	mock.Mock
}

func (svc *PingServiceMock) Ping() string {
	args := svc.Called()

	return args.Get(0).(string)
}
func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	svcMock := new(PingServiceMock)
	svcMock.On("Ping").Return("pong")

	ctrl := NewPingController(svcMock)
	ctrl.Ping(c)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "pong", w.Body.String())
}
