package PUT

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thomas-barth/MyGroceries/db/UPDATE"
)

type JoinRequest struct {
	InviteToken string `json:"invite_token"`
}

func PutJoinList(c *gin.Context) {

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cant get userID"})
		return
	}

	userID := userIDInterface.(uint)

	var req JoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	invite_token := req.InviteToken

	UPDATE.UpdateListWithNewUser(userID, invite_token)

}
