package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userId uint) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("Environment variable JWT_SECRET is not set")
	}
	jwtKey := []byte(secret)
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("Environment variable JWT_SECRET is not set")
	}
	jwtKey := []byte(secret)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, err
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteJSONResponse(w http.ResponseWriter, status int, v any) any {
	w.Header().Add("Content-Type", "application/json")
	return v
}
