package config

import (
	"os"
	"strconv"

	"gorm.io/gorm/logger"
)

var (
	// database connection string
	DSN = ""
	// Database log mode, info = 4, warn = 3, error = 2, silent = 1
	LogMode = logger.Info
)

func Init() {
	DSN = os.Getenv("DSN")

	logLevel, err := strconv.Atoi(os.Getenv("DB_LOG_LEVEL"))
	if err != nil {
		panic("invalid log level")
	}
	LogMode = logger.LogLevel(logLevel)
}

// newrelic license key => e58a81660c32c2a349ec44824c3eab22FFFFNRAL
