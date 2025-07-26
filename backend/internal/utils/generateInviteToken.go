package utils

import (
	"math/rand"
	"thomas-barth/MyGroceries/db/UPDATE"
	"thomas-barth/MyGroceries/internal/models"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateInviteToken(listId uint) string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	var updatedList models.ShoppingList
	updatedList.InviteToken = string(b)
	UPDATE.UpdateListByID(listId, updatedList)

	return string(b)
}
