package http

import (
	"asimov-backend/internal/middleware"
	"net/http"

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
func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	router.ServeHTTP(w, req)
}

func initMiddlewares() {
	authMiddleware := middleware.NewAuthMiddleware()
	router.Use(authMiddleware.Check)
}
