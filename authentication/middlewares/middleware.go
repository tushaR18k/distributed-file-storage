package middlewares

import (
	"authentication/utils"
	"context"
	"fmt"
	"net/http"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Authorization header format invalid", http.StatusUnauthorized)
			return
		}
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		fmt.Println(claims)

		if claims != nil {
			userId := claims.UserId
			ctx := context.WithValue(r.Context(), "userID", userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			fmt.Println("HEYYy!3")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
	})
}
