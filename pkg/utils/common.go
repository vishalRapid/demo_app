package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"strings"
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

// function to create checksum for blog content
func CreateCheckSum(content string) string {
	if len(content) == 0 {
		return content
	}

	h := md5.New()
	io.WriteString(h, content)
	checksum := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Println(checksum)
	return checksum
}

// create slug for blog,
func MakeSlug(title string) (string, error) {

	rand.Seed(time.Now().UTC().UnixNano())
	randomNum := rand.Intn(10000)

	slug := strings.ToLower(title)
	// remove any non-alphanumeric characters
	reg, _ := regexp.Compile("[^a-z0-9]+")
	slug = reg.ReplaceAllString(slug, "-")
	// remove any leading or trailing dashes
	slug = strings.Trim(slug, "-")

	// truncate the slug to the specified length
	if len(slug) > 30 {
		slug = slug[:30]
	}

	// append the random number
	slug += fmt.Sprintf("-%d", randomNum)

	return slug, nil

}
