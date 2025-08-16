package GET

import (
	"net/http"
	"thomas-barth/SharedShopping/db/READ"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	items, err := READ.ReadItems()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "problem with GetItems"})
	}

	c.IndentedJSON(http.StatusOK, items)
}
