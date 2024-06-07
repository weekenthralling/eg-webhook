package storage

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"zjuici.com/tablegpt/eg-webhook/config"
	"zjuici.com/tablegpt/eg-webhook/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {

	var db *gorm.DB
	var err error
	switch cfg.Store.Type {
	case "postgres":
		db, err = gorm.Open(postgres.Open(cfg.Store.Postgres.DSN), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	default:
		log.Fatalf("Unsupported DB type: %s", cfg.Store.Type)
	}

	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to obtain raw database connection: %s", err)
	}

	sqlDB.SetMaxOpenConns(cfg.Store.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Store.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(&models.KernelSession{}); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}

	log.Println("Database initialized successfully")
	return db
}
