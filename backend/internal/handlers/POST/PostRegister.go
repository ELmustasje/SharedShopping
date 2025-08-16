package POST

import (
	"net/http"
	"thomas-barth/SharedShopping/db"
	"thomas-barth/SharedShopping/internal/auth"
	"thomas-barth/SharedShopping/internal/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var creds auth.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	hashedPassword, err := auth.HashPassword(creds.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.User{Username: creds.Username, Password: hashedPassword}
	if result := db.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered"})
}
