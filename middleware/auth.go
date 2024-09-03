package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("secret")

type contextKey string

const userIDContextKey contextKey = "user_id"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token is missing", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.Split(authHeader, " ")[1]
		claims := &jwt.MapClaims{}
		Token, err := jwt.ParseWithClaims(tokenStr, claims, TokenFunc)
		if err != nil || !Token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		//ctx := context.WithValue(r.Context(), "user_id", (*claims)["user_id"])
		ctx := context.WithValue(r.Context(), userIDContextKey, (*claims)["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func TokenFunc(Token *jwt.Token) (interface{}, error) {
	return JwtKey, nil
}
