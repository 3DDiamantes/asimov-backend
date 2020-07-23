package http

import (
	"github.com/gin-gonic/gin"
	"github.com/luxarts/asimov-backend/internal/controller"
	"github.com/luxarts/asimov-backend/internal/service"
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
