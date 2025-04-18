package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	// カスタムクレーム UserID,Username
	// 標準クレーム RegisteredClaims
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// ユーザー名、ユーザーID、シークレットキーを受け取り、JWTトークンを生成する
func GenerateToken(userID uint, username string, secretKey string) (string, error) {
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// トークンの有効期限を24時間に設定
			// トークンの発行時間を現在の時間に設定
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// JWTトークンを生成 署名アルゴリズムはHS256を使用、トークンに含まれる情報はclaims
	// secretKeyを使用して署名されたトークン文字列を返す
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// クライアントから送信されたJWTトークンを検証する関数
func ValidateToken(tokenString string, secretKey string) (*JWTClaims, error) {
	// トークンを解析し、JWTClaims構造体にマッピング
	// 第一引数はトークン文字列、第二引数はJWTClaims構造体のポインタ
	// 第三引数は署名検証のための関数
	// シークレットキーを使用してトークンを検証
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
