package helpers

import (
	"context"
	"encoding/json"
	"time"

	"go-auth/model"

	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func CreateSession(client *redis.Client, user *goth.User) (*model.Session, error) {
	sessionID := generateSessionID() // Generate a unique ID
	session := &model.Session{
		ID:           sessionID,
		UserID:       user.UserID,
		RefreshToken: user.RefreshToken,
		ExpiresAt:    time.Now().Add(time.Hour * 24), // Set expiration time
	}

	sessionData, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}

	err = client.Set(ctx, sessionID, sessionData, session.ExpiresAt.Sub(time.Now())).Err()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func GetSession(client *redis.Client, sessionID string) (*model.Session, error) {
	sessionData, err := client.Get(ctx, sessionID).Result()
	if err == redis.Nil {
		return nil, nil // Session not found
	} else if err != nil {
		return nil, err
	}

	var session model.Session
	err = json.Unmarshal([]byte(sessionData), &session)
	if err != nil {
		return nil, err
	}

	// Check if session is expired
	if session.ExpiresAt.Before(time.Now()) {
		return nil, nil // Session expired
	}

	return &session, nil
}

func DeleteSession(client *redis.Client, sessionID string) error {
	return client.Del(ctx, sessionID).Err()
}

// Helper functions
func generateSessionID() string {
	return uuid.New().String()
}
