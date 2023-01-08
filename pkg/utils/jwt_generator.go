package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateNewAccessToken() (string, error) {
	// Set secret key from .env
	secret := os.Getenv("JWT_SECRET_KEY")

	// Set expires minutes count for secret
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// Create a new claims
	claims := jwt.MapClaims{}

	// Set public claims
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount))

	// Create a new JWT access token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// Generate token
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}