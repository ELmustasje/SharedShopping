package READ

import (
	"thomas-barth/SharedShopping/db"
	"thomas-barth/SharedShopping/internal/models"
)

func ReadUserPassword(userName string) (string, error) {
	var user models.User

	// This is the correct way to query by a named column
	result := db.DB.Where("username = ?", userName).First(&user)

	if result.Error != nil {
		return "", result.Error
	}

	return user.Password, nil
}
