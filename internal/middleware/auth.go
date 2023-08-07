package middleware

import (
	"context"
	"errors"
	"fmt"
	"go-friend-sphere/config"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (m *MiddlewareManager) JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		tokenString := bearerToken[1]
		claims, err := ValidateJWTToken(tokenString, r.Context(), m.cfg)
		if err != nil {
			http.Error(w, "Invalid Authorization token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "userClaims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func ValidateJWTToken(tokenString string, ctx context.Context, c *config.Config) (*jwt.RegisteredClaims, error) {
	signingKey := []byte(c.Server.JwtSecretKey)
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
