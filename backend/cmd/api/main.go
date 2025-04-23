package main

import (
	"english-learning-app/internal/api/routes"
	"english-learning-app/internal/config"
	"english-learning-app/internal/database"
	"english-learning-app/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	// configの初期化
	cfg := config.New()

	// configの読み込み
	gin.SetMode(cfg.GinMode)

	// データベースの接続
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run database migrations
	migrationDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)
	if err := database.RunMigrations(migrationDSN); err != nil {
		log.Printf("Warning: Migration failed: %v", err)
	}

	// ルータの初期化
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://english-learning-app-frontend.onrender.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Health check route
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

	// ルートのセットアップ
	routes.SetupRoutes(r, db)

	// 管理者用の処理
	// 管理者用APIグループの追加（ステップ1）
	admin := r.Group("/admin")

	// 管理者用認証ミドルウェアの追加（ステップ2）
	admin.Use(func(c *gin.Context) {
		// ヘッダーからトークンを取得
		adminToken := c.GetHeader("X-Admin-Token")
		// 環境変数から期待されるトークンを取得
		expectedToken := getEnv("ADMIN_TOKEN", "default_secure_token")

		// トークンが一致しない場合は認証エラー
		if adminToken != expectedToken {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		// 認証成功の場合は次の処理へ
		c.Next()
	})

	// ユーザー一覧取得エンドポイントの実装（ステップ3）
	admin.GET("/users", func(c *gin.Context) {
		var users []models.User
		// データベースからすべてのユーザーを取得
		if err := db.Find(&users).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// パスワードハッシュなどの機密情報を隠す
		for i := range users {
			users[i].PasswordHash = "[HIDDEN]"
		}

		// ユーザーリストをJSON形式で返す
		c.JSON(200, users)
	})

	// Determine port for HTTP service
	port := cfg.Port

	// Start server
	fmt.Printf("Starting server on port %s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// 管理者用トークンを環境変数から取得する関数
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
