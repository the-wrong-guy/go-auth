package model

import "time"

type Session struct {
	ID           string `json:"id"`
	UserID       string `json:"userId"`
	RefreshToken string `json:"refreshToken"`
	// Add other session data as needed
	ExpiresAt time.Time `json:"expiresAt"`
}
