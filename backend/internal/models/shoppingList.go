package models

type ShoppingList struct {
	ID uint `gorm:"primaryKey"`

	Name string `json:"name"`

	// Owner
	OwnerID uint `gorm:"not null"`
	Owner   User `gorm:"foreignKey:OwnerID;references:ID"`

	// Shared users (many to many)
	Users []User `gorm:"many2many:shopping_list_users;"`

	// Items and custom items
	Items       []Item
	CustomItems []CustomItem

	InviteToken string `json:"inviteToken"`
}
