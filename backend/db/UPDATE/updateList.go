package UPDATE

import (
	"thomas-barth/SharedShopping/db"
	"thomas-barth/SharedShopping/internal/models"
)

func UpdateListByID(id uint, list models.ShoppingList) error {
	updates := make(map[string]interface{})

	if list.Name != "" {
		updates["Name"] = list.Name
	}
	if list.CustomItems != nil {
		updates["CustomItems"] = list.CustomItems
	}
	if list.Items != nil {
		updates["Items"] = list.Items
	}
	if list.InviteToken != "" {
		updates["InviteToken"] = list.InviteToken
	}

	return db.DB.Model(&models.ShoppingList{}).Where("id = ?", id).Updates(updates).Error
}
