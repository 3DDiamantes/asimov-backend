package http

import (
	"asimov-backend/internal/controller"
	"asimov-backend/internal/middleware"
	"asimov-backend/internal/service"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func InitRouter() *gin.Engine {
	router = gin.Default()
	initMiddlewares()
	mapRoutes()
	return router
}

func initMiddlewares() {
	authMiddleware := middleware.NewAuthMiddleware()
	router.Use(authMiddleware.Check)
}

func mapRoutes() {
	pingService := service.NewPingService()
	pingController := controller.NewPingController(pingService)
	router.GET("/ping", pingController.Ping)
}
