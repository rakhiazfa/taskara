package services

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/rakhiazfa/gin-boilerplate/constants"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"github.com/rakhiazfa/gin-boilerplate/internal/models"
	"github.com/rakhiazfa/gin-boilerplate/internal/repositories"
	"github.com/rakhiazfa/gin-boilerplate/pkg/oauth"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type OauthService struct {
	db                           *gorm.DB
	userRepository               *repositories.UserRepository
	passwordResetTokenRepository *repositories.PasswordResetTokenRepository
}

func NewOauthService(
	db *gorm.DB,
	userRepository *repositories.UserRepository,
	passwordResetTokenRepository *repositories.PasswordResetTokenRepository,
) *OauthService {
	return &OauthService{
		db:                           db,
		userRepository:               userRepository,
		passwordResetTokenRepository: passwordResetTokenRepository,
	}
}

// Generate google auth url
func (s *OauthService) SignInWithGoogle() string {
	googleOauthConfig := oauth.GetGoogleOauthConfig()

	state := viper.GetString("google_oauth.state")
	url := googleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)

	return url
}

// Handling google oauth callback
func (s *OauthService) GoogleOauthCallback(ctx context.Context, code string, state string) (string, error) {
	// Validate OAuth state
	if state != viper.GetString("google_oauth.state") {
		return "", utils.NewHttpError(http.StatusBadRequest, "Invalid OAuth state", nil)
	}

	googleOauthConfig := oauth.GetGoogleOauthConfig()

	// Exchange authorization code for access token
	token, err := googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return "", utils.NewHttpError(http.StatusInternalServerError, "Failed to get token", err)
	}

	// Fetch user info from Google
	googleUserInfo, err := s.fetchGoogleUserInfo(ctx, &googleOauthConfig, token)
	if err != nil {
		return "", utils.NewHttpError(http.StatusInternalServerError, "Failed to fetch user info", err)
	}

	user, err := s.findOrCreateUser(ctx, googleUserInfo)
	if err != nil {
		return "", err
	}

	if user.Password == "" {
		passwordResetLink, err := s.generatePasswordResetLink(ctx, user)
		if err != nil {
			return "", err
		}

		return passwordResetLink, nil
	}

	return viper.GetString("frontend_url") + "/auth/sign-in", nil
}

// Retrieve user info from Google using OAuth token
func (s *OauthService) fetchGoogleUserInfo(ctx context.Context, config *oauth2.Config, token *oauth2.Token) (*models.GoogleUserInfo, error) {
	client := config.Client(ctx, token)

	res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var userInfo models.GoogleUserInfo
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}

// Find an existing user or create a new user if none is found.
func (s *OauthService) findOrCreateUser(ctx context.Context, googleUserInfo *models.GoogleUserInfo) (*entities.User, error) {
	exsitingUser, err := s.userRepository.FindByEmailUnscoped(googleUserInfo.Email)
	if err != nil {
		return nil, err
	}

	if exsitingUser != nil {
		return exsitingUser, nil
	}

	var user entities.User

	user.Provider = constants.UserProvider_Google

	if err := copier.Copy(&user, &googleUserInfo); err != nil {
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

// Generate a password reset link for the user.
func (s *OauthService) generatePasswordResetLink(ctx context.Context, user *entities.User) (string, error) {
	token, err := s.passwordResetTokenRepository.FindByUserId(user.ID)
	if err != nil {
		return "", err
	}

	// If token doesn't exist or is expired, create or update it
	if token == nil || token.ExpiresAt.Before(time.Now()) {
		token = &entities.PasswordResetToken{
			UserId:    user.ID,
			Token:     utils.GenerateSecureToken(32),
			ExpiresAt: time.Now().Add(15 * time.Minute),
		}

		err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			return s.passwordResetTokenRepository.WithTx(tx).Save(token)
		})
		if err != nil {
			return "", err
		}
	}

	return viper.GetString("frontend_url") + "/auth/reset-password?token=" + token.Token, nil
}
