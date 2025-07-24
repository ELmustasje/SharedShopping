package POST

import (
	"fmt"
	"net/http"
	"thomas-barth/MyGroceries/db/CREATE"
	"thomas-barth/MyGroceries/internal/auth"
	"thomas-barth/MyGroceries/internal/models"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	var newUser models.User
	fmt.Println("called")

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Generate bcrypt hash from the plaintext password
	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	newUser.Password = string(hashedPassword)

	if err := CREATE.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
