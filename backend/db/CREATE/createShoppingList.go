package CREATE

import (
	"fmt"
	"thomas-barth/MyGroceries/db"
	"thomas-barth/MyGroceries/internal/models"
)

func CreateShoppingList(list models.ShoppingList) error {
	if result := db.DB.Create(list); result.Error != nil {
		return result.Error
	}

	fmt.Println("added new list")

	return nil
}

