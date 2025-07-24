package GET

import (
	"net/http"
	"strconv"
	"thomas-barth/MyGroceries/db"
	"thomas-barth/MyGroceries/internal/models"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	id := c.Query("id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	userID, ok := userIDInterface.(uint)
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID type"})
		return
	}

	var list models.ShoppingList
	err = db.DB.Preload("Users").
		Preload("Items").
		Preload("CustomItems").
		First(&list, parsedID).Error

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "list not found"})
		return
	}

	if list.OwnerID == userID {
		c.IndentedJSON(http.StatusOK, list)
		return
	}

	for _, user := range list.Users {
		if user.ID == userID {
			c.IndentedJSON(http.StatusOK, list)
			return
		}
	}

	c.IndentedJSON(http.StatusForbidden, gin.H{"error": "not authorized to view this list"})
}
