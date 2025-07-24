package models

type Item struct {
	ID             uint   `gorm:"primaryKey"`
	Bought         bool   `json:"bought"`
	Name           string `json:"name"`
	Amount         string `json:"amount"` // could also be int if it's always numeric
	ShoppingListID uint   `gorm:"index"`  // foreign key
}
