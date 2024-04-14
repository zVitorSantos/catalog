package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize Postgres
	db, err = InitializePostgreSQL()

	if err != nil {
		return fmt.Errorf("error initializing postgres: %v", err)
	}

	return nil
}

func GetPostgreSQL() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
