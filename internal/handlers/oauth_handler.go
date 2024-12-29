package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakhiazfa/gin-boilerplate/internal/services"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
)

type OauthHandler struct {
	oauthService *services.OauthService
}

func NewOauthHandler(validator *utils.Validator, oauthService *services.OauthService) *OauthHandler {
	return &OauthHandler{
		oauthService: oauthService,
	}
}

func (h *OauthHandler) SignInWithGoogle(c *gin.Context) {
	url := h.oauthService.SignInWithGoogle()

	c.Redirect(http.StatusFound, url)
}

func (h *OauthHandler) GoogleOauthCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	state := c.DefaultQuery("state", "")

	targetUrl, err := h.oauthService.GoogleOauthCallback(c.Request.Context(), code, state)
	utils.PanicIfErr(err)

	c.Redirect(http.StatusFound, targetUrl)
}
