package models

type ShoppingList struct {
	ID uint `gorm:"primaryKey"`

	Name string `json:"name"`

	// Owner
	OwnerID uint
	Owner   User

	// Shared users
	Users []User `gorm:"many2many:shopping_list_users;"`

	// Items and custom items
	Items       []Item
	CustomItems []CustomItem
}
