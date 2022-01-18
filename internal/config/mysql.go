package config

import (
	"fmt"

	"github.com/ffauzann/simpleapi/internal/model/entity"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func connectMySQL(conf *entity.Config) (err error) {
	// Format DSN
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s",
		conf.Database.MySQL.Username,
		conf.Database.MySQL.Password,
		conf.Database.MySQL.Protocol,
		conf.Database.MySQL.Host,
		conf.Database.MySQL.Port,
		conf.Database.MySQL.DBName,
	)

	// Configure connection
	gc := gorm.Config{}
	if conf.App.Debug {
		gc.Logger = logger.Default.LogMode(logger.Info)
	}

	// Open connection
	db, err := gorm.Open(mysql.Open(dsn), &gc)
	if err != nil {
		zap.S().Error(err)
		return
	}

	// set to config
	conf.Connection.MySQL = db

	return
}
