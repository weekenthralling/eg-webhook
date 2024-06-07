package storage

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"time"
	"zjuici.com/tablegpt/eg-webhook/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"zjuici.com/tablegpt/eg-webhook/config"
)

var db *gorm.DB

func Init() {

	cfg := config.GetConfig()
	var err error

	switch cfg.DBType {
	case "postgres":
		db, err = gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	case "mysql":
		db, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	default:
		log.Fatalf("Unsupported DB type: %s", cfg.DBType)
	}

	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to obtain raw database connection: %s", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(&models.KernelSession{}); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}

	log.Println("Database initialized successfully")
}

func GetDB() *gorm.DB {
	return db
}
