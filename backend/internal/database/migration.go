package database

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations runs database migrations from the specified source
func RunMigrations(dbURL string) error {
	// プロジェクトルートからの相対パスでマイグレーションディレクトリを指定
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filepath.Dir(filepath.Dir(b)))
	migrationPath := filepath.Join("file://", basepath, "db", "migrations")

	m, err := migrate.New(
		migrationPath,
		dbURL)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}
