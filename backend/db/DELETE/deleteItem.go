package DELETE

import (
	"thomas-barth/SharedShopping/db"
	"thomas-barth/SharedShopping/internal/models"
)

func DeleteItemsByIDs(ids []uint) error {
	return db.DB.Where("id IN ?", ids).Delete(&models.Item{}).Error
}
