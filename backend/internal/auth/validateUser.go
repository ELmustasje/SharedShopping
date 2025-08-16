package auth

import (
	"thomas-barth/SharedShopping/db/READ"

	"golang.org/x/crypto/bcrypt"
)

func ValidateUser(c Credentials) bool {

	hashpassword, err := READ.ReadUserPassword(c.Username)

	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(c.Password))
	return err == nil
}
