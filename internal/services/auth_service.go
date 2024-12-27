package services

import (
	"github.com/rakhiazfa/gin-boilerplate/internal/models"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (service *AuthService) SignIn(req models.SignInReq) string {
	return ""
}

func (service *AuthService) SignUp(req models.SignUpReq) {
}
