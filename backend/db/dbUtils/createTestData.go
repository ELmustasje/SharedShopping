package dbUtils

import (
	"log"
	"thomas-barth/SharedShopping/db"
	"thomas-barth/SharedShopping/internal/auth"
	"thomas-barth/SharedShopping/internal/models"
)

func CreateTestData() {
	// Passwords
	pw1, _ := auth.HashPassword("alice123")
	pw2, _ := auth.HashPassword("bob123")
	pw3, _ := auth.HashPassword("charlie123")
	pw4, _ := auth.HashPassword("diana123")
	pw5, _ := auth.HashPassword("edward123")

	// Create users (passwords hashed)
	users := []models.User{
		{Username: "alice", Password: pw1},
		{Username: "bob", Password: pw2},
		{Username: "charlie", Password: pw3},
		{Username: "diana", Password: pw4},
		{Username: "edward", Password: pw5},
	}

	if err := db.DB.Create(&users).Error; err != nil {
		log.Fatalf("failed to create users: %v", err)
	}

	// Reload users to access their IDs
	user1 := users[0]
	user2 := users[1]
	user3 := users[2]
	user4 := users[3]
	user5 := users[4]

	// List 1: Shared with Bob and Charlie
	list1 := models.ShoppingList{
		Name:    "Weekly Groceries",
		OwnerID: user1.ID,
		Users:   []models.User{user2, user3},
		Items: []models.Item{
			{Name: "Milk", Amount: "2L", Bought: false},
			{Name: "Eggs", Amount: "12", Bought: true},
			{Name: "Bread", Amount: "1 loaf", Bought: false},
		},
		CustomItems: []models.CustomItem{
			{Name: "Local honey"},
			{Name: "Farm eggs"},
		},
	}

	// List 2: Shared with Alice
	list2 := models.ShoppingList{
		Name:    "Party Supplies",
		OwnerID: user2.ID,
		Users:   []models.User{user1},
		Items: []models.Item{
			{Name: "Chips", Amount: "3 bags", Bought: false},
			{Name: "Soda", Amount: "6 bottles", Bought: false},
			{Name: "Plastic Cups", Amount: "30", Bought: true},
		},
	}

	// List 3: Private list, no shared users
	list3 := models.ShoppingList{
		Name:    "Camping Trip",
		OwnerID: user3.ID,
		Items: []models.Item{
			{Name: "Tent", Amount: "1", Bought: true},
			{Name: "Sleeping Bag", Amount: "2", Bought: false},
			{Name: "Matches", Amount: "3 boxes", Bought: true},
		},
	}

	// List 4: Shared with everyone
	list4 := models.ShoppingList{
		Name:    "Group Vacation",
		OwnerID: user4.ID,
		Users:   []models.User{user1, user2, user3, user5},
		Items: []models.Item{
			{Name: "Sunscreen", Amount: "2 bottles", Bought: false},
			{Name: "Towels", Amount: "5", Bought: false},
		},
	}

	// List 5: Edward’s personal list
	list5 := models.ShoppingList{
		Name:    "Workout Supplies",
		OwnerID: user5.ID,
		Items: []models.Item{
			{Name: "Protein Powder", Amount: "1kg", Bought: false},
			{Name: "Shaker Bottle", Amount: "2", Bought: true},
		},
	}

	if err := db.DB.Create(&[]models.ShoppingList{list1, list2, list3, list4, list5}).Error; err != nil {
		log.Fatalf("failed to create shopping lists: %v", err)
	}

	log.Println("✅ Test data created successfully.")
}
