package config

import (
	"os"

	"github.com/jlordelo/gopportunities.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	// check if database exists
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("data base not found, creating....")
		// create database file and directory
		err := os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// create database and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	// migrate the schemas
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigrate error: %v", err)
		return nil, err
	}
	return db, nil
}
