// backend/internal/services/auth_service_test.go
package services

import (
	"english-learning-app/internal/models"
	"testing"
)

// モックリポジトリの作成
type mockUserRepository struct {
	users map[string]*models.User
}

func (m *mockUserRepository) Create(user *models.User) error {
	m.users[user.Email] = user
	return nil
}

func (m *mockUserRepository) FindByEmail(email string) (*models.User, error) {
	if user, exists := m.users[email]; exists {
		return user, nil
	}
	return nil, nil
}

func (m *mockUserRepository) FindByUsername(username string) (*models.User, error) {
	for _, user := range m.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, nil
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users: make(map[string]*models.User),
	}
}

func TestRegister(t *testing.T) {
	// モックリポジトリの準備
	mockRepo := newMockUserRepository()

	// テスト対象のサービスを作成
	authService := NewAuthService(mockRepo)

	// テストケース1: 新規ユーザー登録が成功するケース
	user, err := authService.Register("testuser", "test@example.com", "password123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user == nil {
		t.Error("Expected user object, got nil")
	}
	if user.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", user.Username)
	}

	// テストケース2: 既存のメールアドレスで登録を試みるケース
	_, err = authService.Register("anotheruser", "test@example.com", "password456")
	if err == nil {
		t.Error("Expected error for duplicate email, got nil")
	}
}

func TestLogin(t *testing.T) {
	// モックリポジトリの準備
	mockRepo := newMockUserRepository()

	// テスト用ユーザーの作成
	testUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
	}
	testUser.SetPassword("password123")
	mockRepo.users[testUser.Email] = testUser

	// テスト対象のサービスを作成
	authService := NewAuthService(mockRepo)

	// テストケース1: 正しい認証情報でログインが成功するケース
	user, err := authService.Login("test@example.com", "password123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user == nil {
		t.Error("Expected user object, got nil")
	}

	// テストケース2: 誤ったパスワードでログインを試みるケース
	_, err = authService.Login("test@example.com", "wrongpassword")
	if err == nil {
		t.Error("Expected error for wrong password, got nil")
	}

	// テストケース3: 存在しないメールアドレスでログインを試みるケース
	_, err = authService.Login("nonexistent@example.com", "password123")
	if err == nil {
		t.Error("Expected error for nonexistent email, got nil")
	}
}
