package UPDATE

import (
	"thomas-barth/MyGroceries/db"
	"thomas-barth/MyGroceries/internal/models"
)

func UpdateItemByID(id uint, item models.Item) error {
	updates := make(map[string]interface{})

	if item.Name != "" {
		updates["Name"] = item.Name
	}
	updates["Bought"] = item.Bought
	if item.Amount != "" {
		updates["Amount"] = item.Amount
	}
	if item.ShoppingListID != 0 {
		updates["ShoppingListID"] = item.ShoppingListID
	}

	return db.DB.Model(&models.Item{}).Where("id = ?", id).Updates(updates).Error
}
