package config

import (
	"gorm.io/gorm/logger"
)

var (
	DSN = "postgres://roshan.r:roshan.r@localhost:5432/change-events"
	// Database log mode, info = 4, warn = 3, error = 2, silent = 1
	LogMode = logger.Info
)

func Init() {
	DSN = "postgres://roshan.r:roshan.r@localhost:5432/change-events" //os.Getenv("DSN")

	// TODO
	// logLevel, err := strconv.Atoi(os.Getenv("DB_LOG_LEVEL"))
	// if err != nil {
	// 	panic("invalid log level")
	// }
	// LogMode = logger.LogLevel(logLevel)
}

// newrelic license key => e58a81660c32c2a349ec44824c3eab22FFFFNRAL
