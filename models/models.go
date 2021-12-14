package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	FirstName  string
	LastName   string
	Email      string
	AWSUserSub string
}
