package models

type CustomItem struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `json:"name"` // free-form text
	ShoppingListID uint   `gorm:"index"`
}
