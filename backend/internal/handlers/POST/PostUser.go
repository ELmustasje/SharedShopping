package POST

import (
	"net/http"
	"thomas-barth/SharedShopping/db/CREATE"
	"thomas-barth/SharedShopping/internal/auth"
	"thomas-barth/SharedShopping/internal/models"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	newUser.Password = hashedPassword

	if err := CREATE.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
