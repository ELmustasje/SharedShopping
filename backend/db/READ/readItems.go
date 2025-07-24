package READ

import (
	"thomas-barth/MyGroceries/db"
	"thomas-barth/MyGroceries/internal/models"
)

func ReadItems() ([]models.Item, error) {
	var items []models.Item
	result := db.DB.Find(&items)

	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}
