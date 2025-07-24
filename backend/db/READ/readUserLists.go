package READ

import (
	"thomas-barth/MyGroceries/db"
	"thomas-barth/MyGroceries/internal/models"
)

func ReadUserLists(userID uint) ([]models.ShoppingList, error) {
	var lists []models.ShoppingList

	err := db.DB.
		Preload("Owner").
		Preload("Users").
		Preload("Items").
		Preload("CustomItems").
		Joins("LEFT JOIN shopping_list_users slu ON shopping_lists.id = slu.shopping_list_id").
		Where("shopping_lists.owner_id = ? OR slu.user_id = ?", userID, userID).
		Group("shopping_lists.id").
		Find(&lists).Error

	if err != nil {
		return nil, err
	}

	return lists, nil
}
