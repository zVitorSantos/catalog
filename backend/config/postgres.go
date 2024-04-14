package config

import (
	"fmt"
	"os"

	"github.com/zVitorSantos/catalog.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Get PostgreSQL details from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, dbname, port)

	// Create DB and connect
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgres connection error: %v", err)
		return nil, err
	}
	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Product{})
	if err != nil {
		logger.Errorf("postgres automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
