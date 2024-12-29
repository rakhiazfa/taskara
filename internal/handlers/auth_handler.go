package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/models"
	"github.com/rakhiazfa/gin-boilerplate/internal/services"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(validator *utils.Validator, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req models.SignUpReq
	utils.PanicIfErr(c.ShouldBind(&req))

	_, err := h.authService.SignUp(c.Request.Context(), &req)
	utils.PanicIfErr(err)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully create a new account",
	})
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var req models.SignInReq
	utils.PanicIfErr(c.ShouldBind(&req))

	token, err := h.authService.SignIn(c.Request.Context(), &req)
	utils.PanicIfErr(err)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req models.ResetPasswordReq
	utils.PanicIfErr(c.ShouldBind(&req))

	err := h.authService.ResetPassword(c.Request.Context(), &req)
	utils.PanicIfErr(err)

	c.JSON(http.StatusOK, gin.H{
		"message": "Reset password successfully",
	})
}
