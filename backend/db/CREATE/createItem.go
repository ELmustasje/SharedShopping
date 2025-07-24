package CREATE

import (
	"fmt"
	"thomas-barth/MyGroceries/db"
	"thomas-barth/MyGroceries/internal/models"
)

func CreateItem(item models.Item) error {
	if result := db.DB.Create(&item); result.Error != nil {
		return result.Error
	}

	fmt.Println("new item created")

	return nil
}
