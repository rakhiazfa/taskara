package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/handlers"
	"github.com/rakhiazfa/gin-boilerplate/internal/middlewares"
)

func InitRoutes(
	authHandler *handlers.AuthHandler,
) *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Recovery())

	publicApi := r.Group("/api")

	initAuthRoutes(publicApi, authHandler)

	return r
}
