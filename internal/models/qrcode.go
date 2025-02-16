package models

import "time"

type QRCode struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:id"`
	Text      string    `json:"text" gorm:"not null"`
	Image     []byte    `json:"image" gorm:"not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
