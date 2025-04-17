package repository

import (
	"english-learning-app/internal/models"
	"errors"

	"gorm.io/gorm"
)

// インターフェイスの定義、メソッドの振る舞いを抽象化
type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
}

// 実装構造体
type userRepository struct {
	db *gorm.DB
}

// 構造体を初期化する関数
// 作成したポインタ型の構造体をインターフェイス型として返す
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// ユーザーを作成する
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// メースアドレスでユーザーを検索する
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// ユーザーを検索する
func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
