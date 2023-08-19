package db

import (
	"sync"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	dbcli  *gorm.DB
	oncedb sync.Once
)

// InitDB to init database
func InitDB(sdn string) (db *gorm.DB, err error) {
	oncedb.Do(func() {
		newLogger := logger.New(
			logrus.NewWriter(), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL Threshold
				IgnoreRecordNotFoundError: true,
				LogLevel:                  logger.Info, // Log level
				Colorful:                  true,        // Disable color printing
			},
		)
		db, err = gorm.Open(postgres.Open(sdn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			PrepareStmt:                              true,
			DisableNestedTransaction:                 true,
			SkipDefaultTransaction:                   true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: viper.GetString("gorm.TablePrefix"), // 表前缀
			},
			Logger: newLogger,
		})
		if err != nil {
			return
		}
		if err = db.Use(tracing.NewPlugin()); err != nil {
			return
		}
		if viper.GetBool("gorm.Debug") {
			db = db.Debug()
		}

		sqlDB, err := db.DB()
		if err != nil {
			return
		}
		sqlDB.SetMaxIdleConns(viper.GetInt("gorm.MaxOpenConns"))                                // 最大空闲连接
		sqlDB.SetMaxOpenConns(viper.GetInt("gorm.MaxIdleConns"))                                // 最大打开连接数
		sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("gorm.MaxLifetime")) * time.Second) // 连接最大存活时间

		dbcli = db
	})
	return dbcli, err
}
