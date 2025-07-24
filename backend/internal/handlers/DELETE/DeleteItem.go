package DELETE

import (
	"net/http"
	"thomas-barth/MyGroceries/db/DELETE"

	"github.com/gin-gonic/gin"
)

func DeleteItemsByIds(c *gin.Context) {
	var ids []uint

	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: expected array of IDs"})
		return
	}

	if len(ids) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No IDs provided"})
		return
	}

	if err := DELETE.DeleteItemsByIDs(ids); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Items deleted", "count": len(ids)})
}

