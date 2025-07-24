package READ

import (
	"thomas-barth/MyGroceries/db"
	"thomas-barth/MyGroceries/internal/models"
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
