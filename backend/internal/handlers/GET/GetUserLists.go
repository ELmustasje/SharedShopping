package GET

import (
	"net/http"
	"thomas-barth/MyGroceries/db/READ"

	"github.com/gin-gonic/gin"
)

func GetUserLists(c *gin.Context) {
	// Get userID from context (set by JWT middleware)
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	// Now call your database logic
	lists, err := READ.ReadUserLists(userID) // convert to uint64 if needed
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shopping lists"})
		return
	}

	c.IndentedJSON(http.StatusOK, lists)
}
