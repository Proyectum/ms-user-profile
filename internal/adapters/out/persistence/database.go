package persistence

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

var logger_levels map[string]logger.LogLevel

func init() {
	logger_levels = map[string]logger.LogLevel{
		"SILENT": logger.Silent,
		"ERROR":  logger.Error,
		"WARN":   logger.Warn,
		"INFO":   logger.Info,
	}
}

func getDatasource() *gorm.DB {

	dbOnce.Do(func() {
		dbUser := viper.GetString("data.datasource.postgres.user")
		host := viper.GetString("data.datasource.postgres.host")
		dbPass := viper.GetString("data.datasource.postgres.password")
		dbName := viper.GetString("data.datasource.postgres.database")
		port := viper.GetInt("data.datasource.postgres.port")
		logLevel := viper.GetString("data.jdbc.gorm.logger.level")

		if len(logLevel) == 0 {
			logLevel = "SILENT"
		}

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", host, dbUser, dbPass, dbName, port)
		var err error
		dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger_levels[logLevel]),
		})
		if err != nil {
			log.Fatal("failed to connect database")
		}
	})

	return dbInstance
}
