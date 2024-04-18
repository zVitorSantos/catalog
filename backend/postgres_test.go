package main

import (
	"fmt"
	"testing"

	"github.com/zVitorSantos/catalog.git/config"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := config.InitializePostgreSQL()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get database: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to the database")

	err = sqlDB.Close()
	if err != nil {
		t.Fatalf("Failed to close database: %v", err)
	}

	fmt.Println("Successfully closed the database connection")
}
