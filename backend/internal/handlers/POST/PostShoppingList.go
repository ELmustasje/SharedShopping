package POST

import (
	"net/http"
	"thomas-barth/MyGroceries/internal/models"

	"github.com/gin-gonic/gin"
)

func PostShoppingList(c *gin.Context) {
	var newShoppingList models.ShoppingList

	if err := c.BindJSON(&newShoppingList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}

}
