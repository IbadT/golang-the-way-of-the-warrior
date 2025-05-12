package db

import (
	"log"

	taskservice "github.com/IbadT/golang-the-way-of-the-warrior.git/internal/taskService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=warrior port=5432 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := DB.AutoMigrate(&taskservice.Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	return DB, nil
}
