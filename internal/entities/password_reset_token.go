package entities

import (
	"time"

	"github.com/google/uuid"
)

type PasswordResetToken struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserId    uuid.UUID
	User      *User     `gorm:"foreignKey:UserId;references:ID"`
	Token     string    `gorm:"type:varchar(100)"`
	ExpiresAt time.Time `gorm:"type:timestamp"`
	CreatedAt time.Time `gorm:"<-:create"`
}
