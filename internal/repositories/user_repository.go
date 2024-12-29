package repositories

import (
	"github.com/google/uuid"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) WithTx(tx *gorm.DB) *UserRepository {
	return &UserRepository{db: tx}
}

func (r *UserRepository) Save(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) FindById(id uuid.UUID) (*entities.User, error) {
	var user entities.User

	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmailUnscoped(email string, exclude ...uuid.UUIDs) (*entities.User, error) {
	var user entities.User

	q := r.db.Unscoped().Where("email = ?", email)

	if len(exclude) > 0 {
		q = q.Not("id IN ?", exclude[0])
	}

	if err := q.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
