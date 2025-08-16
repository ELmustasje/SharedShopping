package CREATE

import (
	"thomas-barth/SharedShopping/db"
	"thomas-barth/SharedShopping/internal/models"
)

func CreateShoppingList(list models.ShoppingList) error {
	if result := db.DB.Create(&list); result.Error != nil {
		return result.Error
	}
	return nil
}
