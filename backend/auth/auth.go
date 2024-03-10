package auth

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"

	"github.com/gorilla/sessions"
	gothGoogle "github.com/markbates/goth/providers/google"
)

func InitAuth() {
	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days
	isProd := false             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd
	store.MaxAge(maxAge)

	gothic.Store = store

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	goth.UseProviders(
		gothGoogle.New(googleClientID, googleClientSecret, "http://localhost:3000/auth/google/callback", "email", "profile"),
	)
}
