package main

import (
	"fmt"
	"thomas-barth/SharedShopping/db"

	"thomas-barth/SharedShopping/internal/auth"
	"thomas-barth/SharedShopping/internal/handlers/GET"
	"thomas-barth/SharedShopping/internal/handlers/POST"
	"thomas-barth/SharedShopping/internal/handlers/PUT"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ No .env file found, falling back to system env")
	}
	fmt.Println("Starting backend")
	// Initialize DB
	db.InitDB()
	//dbUtils.CreateTestData()

	// Setup router
	router := gin.Default()
	configureCORS(router)

	// Public API routes
	api := router.Group("/api")
	{
		api.POST("/postUser", POST.PostUser)
		api.POST("/login", POST.PostLogin)

	}

	// Protected API routes
	protected := api.Group("/")
	protected.Use(auth.JWTAuthMiddleware())
	{
		protected.GET("/protected-endpoint", func(c *gin.Context) {
			userID := c.MustGet("userID").(uint)
			c.JSON(200, gin.H{"msg": "You are authenticated", "user_id": userID})
		})
		protected.GET("/getUserLists", GET.GetUserLists)
		protected.GET("/getList", GET.GetList)
		protected.GET("/getInviteToken", GET.GetInviteToken)

		protected.POST("/postItem", POST.PostItem)
		protected.POST("/postList", POST.PostShoppingList)

		protected.PUT("/putJoin", PUT.PutJoinList)
		protected.PUT("/putItem", PUT.PutItem)
	}

	// Start server
	router.Run("localhost:8080")
}

func configureCORS(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
