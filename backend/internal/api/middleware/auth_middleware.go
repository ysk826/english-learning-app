package middleware

import (
	"english-learning-app/internal/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware は認証が必要なルートへのアクセス制御を行うミドルウェア関数
// JWTトークンの検証を行い、有効な場合はリクエストを続行し、無効な場合はエラーを返す
//
// 処理内容:
// 1. 'Authorization' ヘッダーからBearerトークンを取得
// 2. JWTトークンの検証
// 3. 有効な場合、ユーザー情報をコンテキストに格納し後続の処理に利用可能にする
// 4. 無効な場合、401 Unauthorizedエラーを返す
//
// 使用例:
//
//	protected := router.Group("/api")
//	protected.Use(middleware.AuthMiddleware())
//	{
//	  protected.GET("/profile", profileHandler)
//	}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}

		// トークンを分割してBearerトークンを取得
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
			c.Abort()
			return
		}

		// トークンを取得
		// parts[0]は"Bearer"、parts[1]はトークン
		tokenString := parts[1]

		// シークレットキーを使用
		// todo: JWT_SECRETを環境変数から取得する
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			jwtSecret = "your-default-secret-key"
		}

		// トークンを検証
		claims, err := utils.ValidateToken(tokenString, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// トークンが有効な場合、ユーザー情報をコンテキストに格納
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
