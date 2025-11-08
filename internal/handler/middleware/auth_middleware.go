package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"users/pkgs/auth"
)

type contextKey string

const UserCtxKey contextKey = "user"

type AuthMiddleware interface {
	AuthenticateToken(next http.Handler) http.Handler
}

type jwtAuthMiddleware struct{}

func NewAuthMiddleware() AuthMiddleware { return &jwtAuthMiddleware{} }

func (m *jwtAuthMiddleware) AuthenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]
		claims, err := auth.ValidateToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Attach user info to context
		ctx := context.WithValue(r.Context(), UserCtxKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type UserInfo struct {
	UserID string
	Email  string
}

func GetUserFromContext(ctx context.Context) (*UserInfo, error) {
	claims, ok := ctx.Value(UserCtxKey).(*auth.Claims)
	if !ok || claims == nil {
		return nil, errors.New("user not found in context")
	}

	return &UserInfo{
		UserID: claims.UserID,
		Email:  claims.Email,
	}, nil
}
