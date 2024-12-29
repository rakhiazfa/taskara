package repositories

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
	"gorm.io/gorm"
)

type PasswordResetTokenRepository struct {
	db *gorm.DB
}

func NewPasswordResetTokenRepository(db *gorm.DB) *PasswordResetTokenRepository {
	return &PasswordResetTokenRepository{db: db}
}

func (r *PasswordResetTokenRepository) WithTx(tx *gorm.DB) *PasswordResetTokenRepository {
	return &PasswordResetTokenRepository{db: tx}
}

func (r *PasswordResetTokenRepository) Save(passwordResetToken *entities.PasswordResetToken) error {
	return r.db.Save(passwordResetToken).Error
}

func (r *PasswordResetTokenRepository) FindByUserId(userId uuid.UUID) (*entities.PasswordResetToken, error) {
	var passwordResetToken entities.PasswordResetToken

	if err := r.db.Where("user_id = ?", userId).First(&passwordResetToken).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &passwordResetToken, nil
}

func (r *PasswordResetTokenRepository) FindByToken(token string) (*entities.PasswordResetToken, error) {
	var passwordResetToken entities.PasswordResetToken

	if err := r.db.Where("token = ?", token).First(&passwordResetToken).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &passwordResetToken, nil
}

func (r *PasswordResetTokenRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&entities.PasswordResetToken{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return utils.NewHttpError(http.StatusNotFound, "Password reset token not found", nil)
	}

	return nil
}
