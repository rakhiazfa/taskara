package services

import (
	"context"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/gin-boilerplate/constants"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"github.com/rakhiazfa/gin-boilerplate/internal/models"
	"github.com/rakhiazfa/gin-boilerplate/internal/repositories"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	db                           *gorm.DB
	validator                    *utils.Validator
	userRepository               *repositories.UserRepository
	passwordResetTokenRepository *repositories.PasswordResetTokenRepository
}

func NewAuthService(
	db *gorm.DB,
	validator *utils.Validator,
	userRepository *repositories.UserRepository,
	passwordResetTokenRepository *repositories.PasswordResetTokenRepository,
) *AuthService {
	return &AuthService{
		db:                           db,
		validator:                    validator,
		userRepository:               userRepository,
		passwordResetTokenRepository: passwordResetTokenRepository,
	}
}

func (s *AuthService) SignIn(ctx context.Context, req *models.SignInReq) (string, error) {
	if err := s.validator.Validate(req); err != nil {
		return "", err
	}

	return "", nil
}

func (s *AuthService) SignUp(ctx context.Context, req *models.SignUpReq) (*entities.User, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	userWithSameEmail, err := s.userRepository.FindByEmailUnscoped(req.Email)
	if err != nil {
		return nil, err
	}

	if userWithSameEmail != nil {
		return nil, utils.NewHttpError(http.StatusConflict, "An account with this email already exists", nil)
	}

	var user entities.User
	user.Provider = constants.UserProvider_Local

	if err := copier.Copy(&user, &req); err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return s.userRepository.WithTx(tx).Save(&user)
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) ResetPassword(ctx context.Context, req *models.ResetPasswordReq) error {
	if err := s.validator.Validate(req); err != nil {
		return err
	}

	token, err := s.passwordResetTokenRepository.FindByToken(req.Token)
	if err != nil {
		return err
	}

	if token == nil || token.ExpiresAt.Before(time.Now()) {
		return utils.NewHttpError(http.StatusBadRequest, "Invalid password reset token", nil)
	}

	user, err := s.userRepository.FindById(token.UserId)
	if err != nil {
		return err
	}
	if user == nil {
		return utils.NewHttpError(http.StatusBadRequest, "User not found", nil)
	}

	hash, err := user.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user.Password = hash

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := s.passwordResetTokenRepository.Delete(token.ID)
		if err != nil {
			return err
		}

		return s.userRepository.WithTx(tx).Save(user)
	})
	if err != nil {
		return err
	}

	return nil
}
