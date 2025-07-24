package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"-"`

	// Owned lists
	OwnedLists []ShoppingList `gorm:"foreignKey:OwnerID"`

	// Shared lists
	ShoppingLists []ShoppingList `gorm:"many2many:shopping_list_users;"`
}
