package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luxarts/asimov-backend/internal/service"
)

type PingController interface {
	Ping(c *gin.Context)
}
type pingController struct {
	svc service.PingService
}

func NewPingController(svc service.PingService) PingController {
	return &pingController{svc: svc}
}

func (ctrl *pingController) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
