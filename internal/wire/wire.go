//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/rakhiazfa/gin-boilerplate/internal/handlers"
	"github.com/rakhiazfa/gin-boilerplate/internal/infrastructures"
	"github.com/rakhiazfa/gin-boilerplate/internal/repositories"
	"github.com/rakhiazfa/gin-boilerplate/internal/services"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
	"github.com/rakhiazfa/gin-boilerplate/routes"
)

var userModule = wire.NewSet(
	repositories.NewUserRepository,
)

var oauthModule = wire.NewSet(
	services.NewOauthService,
	handlers.NewOauthHandler,
)

var authModule = wire.NewSet(
	services.NewAuthService,
	handlers.NewAuthHandler,
)

func NewApplication() *gin.Engine {
	wire.Build(
		infrastructures.NewPostgresConnection,
		utils.NewValidator,
		repositories.NewPasswordResetTokenRepository,
		userModule,
		oauthModule,
		authModule,
		routes.InitRoutes,
	)

	return nil
}
