package entities

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseEntityWithSoftDelete
	Provider       string `gorm:"type:varchar(100);default:null"`
	GoogleId       string `gorm:"type:varchar(100);default:null"`
	ProfilePicture string `gorm:"type:varchar(100);default:null"`
	Name           string `gorm:"type:varchar(100)"`
	Email          string `gorm:"type:varchar(100);unique"`
	Password       string `gorm:"type:varchar(100);default:null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := u.HashPassword(u.Password)
		if err != nil {
			return err
		}

		u.Password = hash
	}

	return
}

func (u *User) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
