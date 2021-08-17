package controller

import (
	"net/http"

	"asimov-backend/internal/service"

	"github.com/gin-gonic/gin"
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
	response := ctrl.svc.Ping()

	c.String(http.StatusOK, response)
}
