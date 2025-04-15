package main

import (
	"fmt"
	"log"

	"english-learning-app/internal/config"
	"english-learning-app/internal/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	// Load configuration
	cfg := config.New()

	// Set Gin mode
	gin.SetMode(cfg.GinMode)

	// Initialize database connection
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize router
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Routes
	r.GET("/health", func(c *gin.Context) {
		// Check database connection
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Database connection error",
			})
			return
		}

		// Ping database
		if err := sqlDB.Ping(); err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Database ping failed",
			})
			return
		}

		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "API server is running and connected to the database",
		})
	})

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		// These will be implemented later
		auth := v1.Group("/auth")
		{
			auth.POST("/register", func(c *gin.Context) {
				c.JSON(501, gin.H{"message": "Not implemented yet"})
			})
			auth.POST("/login", func(c *gin.Context) {
				c.JSON(501, gin.H{"message": "Not implemented yet"})
			})
		}

		words := v1.Group("/words")
		{
			words.GET("", func(c *gin.Context) {
				c.JSON(501, gin.H{"message": "Not implemented yet"})
			})
			words.POST("", func(c *gin.Context) {
				c.JSON(501, gin.H{"message": "Not implemented yet"})
			})
		}

		sentences := v1.Group("/sentences")
		{
			sentences.GET("", func(c *gin.Context) {
				c.JSON(501, gin.H{"message": "Not implemented yet"})
			})
			sentences.POST("", func(c *gin.Context) {
				c.JSON(501, gin.H{"message": "Not implemented yet"})
			})
		}
	}

	// Determine port for HTTP service
	port := cfg.Port

	// Start server
	fmt.Printf("Starting server on port %s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
