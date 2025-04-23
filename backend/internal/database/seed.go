package database

import (
	"fmt"
	"io/ioutil"
	"log"

	"gorm.io/gorm"
)

// RunSeeds runs seed data from the specified file
func RunSeeds(db *gorm.DB, seedFile string) error {
	log.Printf("Running seed file: %s", seedFile)

	// シードファイルの読み込み
	seedData, err := ioutil.ReadFile(seedFile)
	if err != nil {
		return fmt.Errorf("failed to read seed file: %w", err)
	}

	// シードデータの実行
	result := db.Exec(string(seedData))
	if result.Error != nil {
		return fmt.Errorf("failed to execute seed SQL: %w", result.Error)
	}

	log.Println("Seed data loaded successfully")
	return nil
}
