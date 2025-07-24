package PUT

import (
	"fmt"
	"net/http"
	"thomas-barth/MyGroceries/db/UPDATE"
	"thomas-barth/MyGroceries/internal/models"

	"github.com/gin-gonic/gin"
)

func PutItem(c *gin.Context) {
	var updatedItem models.Item
	fmt.Println("called")

	// Parse JSON body
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := UPDATE.UpdateItemByID(updatedItem.ID, updatedItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update item"})
	}

}
