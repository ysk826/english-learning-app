package routes

import (
	"english-learning-app/internal/api/handlers"
	"english-learning-app/internal/api/middleware"
	"english-learning-app/internal/repository"
	"english-learning-app/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Public routes
	v1 := r.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes
		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// Words routes (to be implemented)
			words := protected.Group("/words")
			{
				words.GET("", func(c *gin.Context) {
					c.JSON(501, gin.H{"message": "Not implemented yet"})
				})
				words.POST("", func(c *gin.Context) {
					c.JSON(501, gin.H{"message": "Not implemented yet"})
				})
			}

			// Sentences routes (to be implemented)
			sentences := protected.Group("/sentences")
			{
				sentences.GET("", func(c *gin.Context) {
					c.JSON(501, gin.H{"message": "Not implemented yet"})
				})
				sentences.POST("", func(c *gin.Context) {
					c.JSON(501, gin.H{"message": "Not implemented yet"})
				})
			}
		}
	}
}
