package http

import (
	"asimov-backend/internal/controller"
	"asimov-backend/internal/service"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func InitRouter() *gin.Engine {
	router = gin.Default()
	MapRoutes()
	return router
}

func MapRoutes() {
	pingService := service.NewPingService()
	pingController := controller.NewPingController(pingService)
	router.GET("/ping", pingController.Ping)
}
