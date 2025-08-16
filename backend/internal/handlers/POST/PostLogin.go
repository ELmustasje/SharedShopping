package POST

import (
	"net/http"
	"thomas-barth/SharedShopping/db"
	"thomas-barth/SharedShopping/internal/auth"
	"thomas-barth/SharedShopping/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostLogin(c *gin.Context) {
	var credentials auth.Credentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

	var user models.User
	if err := db.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	if !auth.ValidateUser(credentials) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect password"})
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}
