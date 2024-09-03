package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	GormDB *gorm.DB
}

func SetupDB() *DB {
	dsn := "host=localhost user=postgres password=momdad143 dbname=bookdata port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&Book{}, &User{})
	if err != nil {
		log.Fatalf("Failed to migrate schema: %v", err)
		return nil
	}

	return &DB{GormDB: db}
}
