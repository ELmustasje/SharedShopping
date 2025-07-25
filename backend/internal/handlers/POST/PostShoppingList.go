package POST

import (
	"net/http"
	"thomas-barth/MyGroceries/db/CREATE"
	"thomas-barth/MyGroceries/internal/models"

	"github.com/gin-gonic/gin"
)

func PostShoppingList(c *gin.Context) {
	var newShoppingList models.ShoppingList

	if err := c.BindJSON(&newShoppingList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	ownerIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find user"})
	}

	ownerID := ownerIDInterface.(uint)

	newShoppingList.OwnerID = ownerID

	if err := CREATE.CreateShoppingList(newShoppingList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save ShoppingList"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ShoppingList created successfully"})

}
