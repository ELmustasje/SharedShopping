package CREATE

import (
	"fmt"
	"thomas-barth/MyGroceries/db"
	"thomas-barth/MyGroceries/internal/models"
)

func CreateUser(user models.User) error {
	if result := db.DB.Create(&user); result.Error != nil {
		return result.Error
	}

	fmt.Println("new user created")

	return nil
}
