package auth

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

//	func HashPassword(password string) (string, error) {
//		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//		return string(bytes), err
//	}
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}
	return string(hashed), err
}
