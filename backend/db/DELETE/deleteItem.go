package DELETE

import (
	"thomas-barth/MyGroceries/db"
	"thomas-barth/MyGroceries/internal/models"
)

func DeleteItemsByIDs(ids []uint) error {
	return db.DB.Where("id IN ?", ids).Delete(&models.Item{}).Error
}
