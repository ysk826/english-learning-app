package services

import (
	"english-learning-app/internal/models"
	"english-learning-app/internal/repository"
	"errors"
)

// インターフェイスの定義、メソッドの振る舞いを抽象化
type AuthService interface {
	Register(username, email, password string) (*models.User, error)
	Login(email, password string) (*models.User, error)
}

// 実装構造体
type authService struct {
	userRepo repository.UserRepository
}

// 構造体を初期化する関数
// 作成したポインタ型の構造体をインターフェイス型として返す
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// ユーザー登録
func (s *authService) Register(username, email, password string) (*models.User, error) {
	// メールアドレスを使用してユーザーを検索
	existingUser, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("このメールアドレスはすでに使用されています")
	}

	// ユーザー名の重複チェック
	existingUser, err = s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}

	// 新しいユーザーを作成
	user := &models.User{
		Username: username,
		Email:    email,
	}

	// パスワードをハッシュ化
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	// ユーザーをデータベースに保存
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// ログイン処理
func (s *authService) Login(email, password string) (*models.User, error) {
	// メールアドレスを使用してユーザーを検索
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// パスワードを検証
	if !user.CheckPassword(password) {
		return nil, errors.New("invalid credentials")
	}

	// ログインが成功した場合、ユーザー情報を返す
	return user, nil
}
