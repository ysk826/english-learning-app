package routes

import (
	"english-learning-app/internal/api/handlers"
	"english-learning-app/internal/api/middleware"
	"english-learning-app/internal/repository"
	"english-learning-app/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ルートとapi設定を行う
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// データベース接続（db）を外部から受け取り、それをリポジトリに注入
	userRepo := repository.NewUserRepository(db)

	// リポジトリをサービスに注入
	authService := services.NewAuthService(userRepo)

	// サービスをハンドラに注入
	authHandler := handlers.NewAuthHandler(authService)

	// api/v1 ルートグループを作成
	v1 := r.Group("/api/v1")
	{
		// サブグループを作成
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// サブグループを作成、親グループのパスを継承
		// protectedは認証ミドルウェアを使用
		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// wordsグループを作成
			words := protected.Group("/words")
			{
				// todo ユーザーの単語、文章リストを取得する処理、新しい単語、文章をデータベースに追加する処理
				words.GET("", func(c *gin.Context) {
					c.JSON(501, gin.H{"message": "Not implemented yet"})
				})
				words.POST("", func(c *gin.Context) {
					c.JSON(501, gin.H{"message": "Not implemented yet"})
				})
			}

			// sentencesグループを作成
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
