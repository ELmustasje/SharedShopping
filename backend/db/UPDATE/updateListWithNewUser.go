package UPDATE

import (
	"thomas-barth/SharedShopping/db"
	"thomas-barth/SharedShopping/internal/models"
)

func UpdateListWithNewUser(userID uint, inviteToken string) error {
	var list models.ShoppingList
	if err := db.DB.Where("invite_token = ?", inviteToken).First(&list).Error; err != nil {
		return err
	}

	user := &models.User{ID: userID}

	return db.DB.Model(&list).Association("Users").Append(user)
}
