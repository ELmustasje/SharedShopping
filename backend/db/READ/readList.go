package READ

import (
	"thomas-barth/MyGroceries/db"
	"thomas-barth/MyGroceries/internal/models"
)

func ReadList(id uint) (*models.ShoppingList, error) {
	var list models.ShoppingList

	result := db.DB.
		Preload("Items").
		Preload("CustomItems").
		Preload("Owner").
		Preload("Users").
		First(&list, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &list, nil
}
