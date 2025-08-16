package PUT

import (
	"net/http"
	"thomas-barth/SharedShopping/db/UPDATE"
	"thomas-barth/SharedShopping/internal/models"

	"github.com/gin-gonic/gin"
)

func PutItem(c *gin.Context) {
	var updatedItem models.Item

	// Parse JSON body
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := UPDATE.UpdateItemByID(updatedItem.ID, updatedItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update item"})
	}

}
