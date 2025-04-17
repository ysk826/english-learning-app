package handlers

import (
	"english-learning-app/internal/services"
	"english-learning-app/internal/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// 構造体を定義
// AuthServiceに依存する
type AuthHandler struct {
	authService services.AuthService
}

// 構造体を初期化する関数
// 作成したポインタ型の構造体を返す
func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// DTOとしてリクエストボディの構造体を定義
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// DTOとしてリクエストボディの構造体を定義
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// DTOとしてレスポンスボディの構造体を定義
type AuthResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// 登録処理を行う
// ユーザー名、メールアドレス、パスワードを受け取り、JSON形式でレスポンスを返す
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユーザー名、メールアドレス、パスワードを使用してユーザーを登録
	user, err := h.authService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// シークレットキーを使用
	// todo: JWT_SECRETを環境変数から取得する
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-default-secret-key"
	}

	// JWTトークンを生成
	token, err := utils.GenerateToken(user.ID, user.Username, jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// ユーザー名、メールアドレス、トークンを含むJSON形式のレスポンスを返す
	c.JSON(http.StatusCreated, AuthResponse{
		Token:    token,
		Username: user.Username,
		Email:    user.Email,
	})
}

// ログイン処理を行う
// メールアドレス、パスワードを受け取り、JSON形式でレスポンスを返す
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// メールアドレス、パスワードを使用してユーザーをログイン
	user, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// シークレットキーを使用
	// todo: JWT_SECRETを環境変数から取得する
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-default-secret-key"
	}

	// JWTトークンを生成
	token, err := utils.GenerateToken(user.ID, user.Username, jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// ユーザー名、メールアドレス、トークンを含むJSON形式のレスポンスを返す
	c.JSON(http.StatusOK, AuthResponse{
		Token:    token,
		Username: user.Username,
		Email:    user.Email,
	})
}
