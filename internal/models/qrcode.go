package models

import "time"

type QRCode struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Image     string    `json:"image"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
