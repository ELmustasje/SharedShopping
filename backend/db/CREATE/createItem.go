package CREATE

import (
	"thomas-barth/SharedShopping/db"
	"thomas-barth/SharedShopping/internal/models"
)

func CreateItem(item models.Item) error {
	if result := db.DB.Create(&item); result.Error != nil {
		return result.Error
	}

	return nil
}
