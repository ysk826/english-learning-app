package main

import (
	"english-learning-app/internal/config"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// コマンドライン引数の確認
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/migrate/main.go [up|down|version|create NAME]")
	}

	// 設定の読み込み
	cfg := config.New()

	// マイグレーションDSNの設定
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)

	// コマンドの処理
	command := os.Args[1]

	// create コマンドの処理
	if command == "create" && len(os.Args) >= 3 {
		// 新しいマイグレーションファイルを作成するロジック
		// ここでは外部コマンドの migrate を使用する
		log.Printf("Creating new migration: %s", os.Args[2])
		log.Println("Run: migrate create -ext sql -dir db/migrations -seq", os.Args[2])
		return
	}

	// マイグレーションインスタンスの作成
	m, err := migrate.New(
		"file://db/migrations",
		dsn)
	if err != nil {
		log.Fatal(err)
	}

	// コマンドに応じた処理
	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration up completed successfully")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration down completed successfully")
	case "version":
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Current version: %d, Dirty: %v", version, dirty)
	default:
		log.Fatal("Unknown command. Use: up, down, version, or create")
	}
}
