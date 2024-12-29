package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/handlers"
)

func initOauthRoutes(r *gin.RouterGroup, handler *handlers.OauthHandler) {
	group := r.Group("/oauth")

	group.GET("/sign-in/google", handler.SignInWithGoogle)
	group.GET("/callback/google", handler.GoogleOauthCallback)
}
