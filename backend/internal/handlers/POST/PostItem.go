package POST

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"thomas-barth/MyGroceries/db/CREATE"
	"thomas-barth/MyGroceries/internal/models"
)

func PostItem(c *gin.Context) {
	var newItem models.Item

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}

	listsInterface, exists := c.Get("listIDs")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "No access list found in context"})
		return
	}

	listIDs, ok := listsInterface.([]uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid list ID format in context"})
		return
	}

	allowed := slices.Contains(listIDs, newItem.ShoppingListID)

	if !allowed {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to modify this list"})
		return
	}

	if err := CREATE.CreateItem(newItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save item"})
	}
}
