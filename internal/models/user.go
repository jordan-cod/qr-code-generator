package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:id"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	Email     string    `json:"email" gorm:"size:255;unique;not null"`
	Password  string    `json:"password" gorm:"size:255;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updatedAt;autoUpdateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
