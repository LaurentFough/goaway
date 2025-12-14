package database

import (
	"log"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Initialize() *gorm.DB {
	databasePath, err := GetDatabasePath()
	if err != nil {
		log.Fatal("failed to determine database path: %v", err)
	}
	
	if err := os.MkdirAll(filepath.Dir(databasePath), 0755); err != nil {
		log.Fatal("failed to create data directory: %v", err)
	}
	
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed while initializing database: %v", err)
	}
	
	if err := AutoMigrate(db); err != nil {
		log.Fatal("auto migrate failed: %v", err)
	}
	
	return db
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Source{},
		&Blacklist{},
		&Whitelist{},
		&RequestLog{},
		&RequestLogIP{},
		&Resolution{},
		&MacAddress{},
		&User{},
		&APIKey{},
		&Notification{},
		&Prefetch{},
		&Audit{},
		&Alert{},
	)
}
