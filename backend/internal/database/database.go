package database

import (
	"fmt"
	"log"
	"time"

	"english-learning-app/internal/config"
	"english-learning-app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// New creates a new database connection
func New(cfg *config.Config) (*gorm.DB, error) {
	// Configure GORM logger
	gormLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Connect to database with retry logic
	var db *gorm.DB
	var err error
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{
			Logger: gormLogger,
		})

		if err == nil {
			break
		}

		log.Printf("Failed to connect to database, retrying in 5 seconds... (%d/%d)", i+1, maxRetries)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
	}

	// Get the underlying *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Auto migrate the schema
	err = db.AutoMigrate(
		&models.User{},
		&models.Word{},
		&models.Sentence{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate schema: %w", err)
	}

	log.Println("Connected to database successfully")
	return db, nil
}
