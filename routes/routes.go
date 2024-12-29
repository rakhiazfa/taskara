package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/handlers"
	"github.com/rakhiazfa/gin-boilerplate/internal/middlewares"
)

func InitRoutes(
	oauthHandler *handlers.OauthHandler,
	authHandler *handlers.AuthHandler,
) *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Recovery())

	web := r.Group("")
	api := r.Group("/api")

	initOauthRoutes(web, oauthHandler)
	initAuthRoutes(api, authHandler)

	return r
}
