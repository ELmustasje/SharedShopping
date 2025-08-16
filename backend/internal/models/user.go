package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:255;uniqueIndex" json:"username"`
	Password string `gorm:"size:255" json:"password"`

	// Owned lists
	OwnedLists []ShoppingList `gorm:"foreignKey:OwnerID"`

	// Shared lists
	ShoppingLists []ShoppingList `gorm:"many2many:shopping_list_users;"`
}
