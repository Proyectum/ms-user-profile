package persistence

import (
	"fmt"
	"github.com/proyectum/ms-user-profile/internal/boot"
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
		postgresProps := boot.CONFIG.Data.Datasource.Postgres
		jdbcProps := boot.CONFIG.Data.JDBC
		logLevel := jdbcProps.Gorm.Logger.Level

		if len(logLevel) == 0 {
			logLevel = "SILENT"
		}

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
			postgresProps.Host, postgresProps.User, postgresProps.Password, postgresProps.Database, postgresProps.Port)
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
