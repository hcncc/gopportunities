package config

import (
	"os"

	"github.com/hcncc/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("SQLITE")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database is not exists, creating...")

		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("SQLITE Opening erro: %v", err)

		return nil, err
	}

	// Migrate the Schemas
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("SQLITE automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
