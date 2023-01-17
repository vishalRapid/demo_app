package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"golang.org/x/crypto/bcrypt"
)

// hash current password
func HashPassword(password string) string {
	// Hashing the password with a default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hashed password:", string(hashedPassword))
	return string(hashedPassword)
}

// function to compare password if it matches or not
func ComparePassword(hashedPassword string, newPassword string) bool {

	// Comparing the password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(newPassword)); err != nil {
		return false
	} else {
		return true
	}
}

// generate hash
func GenerateJWT(userID string) string {
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := os.Getenv("JWT_SECRET")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString

}

// validate token for user
func ValidateToken(tokenString string) (string, error) {
	// Parse the token using the secret key
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	// Check if the token is valid
	if err != nil {
		return "", fmt.Errorf(constant.TOKEN_MALFORMED)
	} else if !token.Valid {
		return "", fmt.Errorf(constant.TOKEN_MALFORMED)
	} else {
		// Token is valid, get the claims
		claims := token.Claims.(jwt.MapClaims)
		return claims["user_id"].(string), nil
	}
}
