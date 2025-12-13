package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/MananLedwani/taskmanager/backend/internal/utils"
)

type contextKey string

const UserIDKey contextKey = "userID"

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			http.Error(w, "Invalid auth header format", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ParseJWT(parts[1])
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		userID := claims["user_id"].(string)

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
