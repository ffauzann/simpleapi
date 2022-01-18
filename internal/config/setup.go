package config

import (
	"github.com/ffauzann/simpleapi/internal/model/entity"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

func Setup() (conf entity.Config, err error) {
	// setup logger for debugging purpose
	setupLogger()

	// read config/config.yaml file
	err = readConfig(&conf)
	if err != nil {
		zap.S().Error(err)
		return
	}

	// connect to mysql
	err = connectMySQL(&conf)
	if err != nil {
		zap.S().Error(err)
		return
	}

	// Init go validator
	conf.App.Validator = validator.New()

	return
}
