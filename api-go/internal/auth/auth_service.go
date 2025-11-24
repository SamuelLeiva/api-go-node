package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type AuthService struct{}

func (s *AuthService) Login(username, password string) (string, error) {
	// Login “falso” de prueba
	if username != "admin" || password != "123456" {
		return "", ErrInvalidCredentials
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret"
	}

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret"
	}

	return jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		},
	)
}
