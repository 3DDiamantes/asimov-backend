package http

import (
	"asimov-backend/internal/controller"
	"asimov-backend/internal/defines"
	"asimov-backend/internal/service"
)

func mapRoutes() {
	pingService := service.NewPingService()
	pingController := controller.NewPingController(pingService)
	router.GET(defines.PathPingGet, pingController.Ping)
}
