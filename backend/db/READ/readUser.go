package READ

import (
	"thomas-barth/SharedShopping/db"
	"thomas-barth/SharedShopping/internal/models"
)

func ReadUser(userID uint) (*models.User, error) {
	var user models.User

	result := db.DB.First(&user, userID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
