package db

import (
	"github.com/roshan8/change-events/api"
	"github.com/roshan8/change-events/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var store *gorm.DB

func InitializeDB() *gorm.DB {
	// dsn := "appuser:apppass@tcp(mysql)/appdb?charset=utf8mb4&parseTime=True&loc=Local"   // for docker with mysql
	// dsn := "appuser:apppass@tcp(127.0.0.1:3306)/appdb?charset=utf8mb4&parseTime=True&loc=Local" // for local with mysql
	// "host=localhost user=appuser password=apppass dbname=appdb port=5432 sslmode=disable TimeZone=Asia/Kolkata" // for local pg
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(config.LogMode),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize the database")
	}

	// initialize the global store object
	store = db

	if err = db.Migrator().AutoMigrate(&api.KubernetesEvent{}); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate the database")
	}

	log.Info().Msg("db is inited")
	return db
}
