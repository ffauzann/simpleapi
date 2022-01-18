package entity

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type (
	Config struct {
		App        ConfigApp      `yaml:"app"`
		Database   ConfigDatabase `yaml:"database"`
		Connection ConfigConnection
	}

	ConfigApp struct {
		Host      string    `yaml:"host"`
		Port      int       `yaml:"port"`
		Debug     bool      `yaml:"debug"`
		JWT       ConfigJWT `yaml:"jwt"`
		Validator *validator.Validate
	}

	ConfigJWT struct {
		Secret string `yaml:"secret"`
	}

	ConfigDatabase struct {
		MySQL ConfigURI `yaml:"mysql"`
	}

	ConfigConnection struct {
		MySQL *gorm.DB
	}

	ConfigURI struct {
		Protocol string `yaml:"protocol"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	}
)
