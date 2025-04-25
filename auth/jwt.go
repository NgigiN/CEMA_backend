package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// CreateJWT generates a JWT token for a given email.
// It takes a secret key and an email as input and returns the signed token or an error.
func CreateJWT(secret []byte, email string) (string, error) {
	// Define the token claims
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
