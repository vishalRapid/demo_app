package utils

import (
	"fmt"

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
