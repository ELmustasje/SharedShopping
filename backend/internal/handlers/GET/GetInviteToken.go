package GET

import (
	"net/http"
	"strconv"
	"thomas-barth/SharedShopping/db/READ"
	"thomas-barth/SharedShopping/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetInviteToken(c *gin.Context) {
	rawListId := c.Query("listId")
	listId, err := strconv.Atoi(rawListId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	list, err := READ.ReadList(uint(listId))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "read list"})
	}

	if list.InviteToken == "" {
		c.IndentedJSON(http.StatusOK, utils.GenerateInviteToken(uint(listId)))
		return
	}

	c.IndentedJSON(http.StatusOK, list.InviteToken)
}
