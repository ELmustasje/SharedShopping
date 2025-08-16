package CREATE

import (
	"thomas-barth/SharedShopping/db"
	"thomas-barth/SharedShopping/internal/models"
)

func CreateUser(user models.User) error {
	if result := db.DB.Create(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
