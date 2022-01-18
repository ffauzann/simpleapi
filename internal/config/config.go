package config

import (
	"io/ioutil"

	"github.com/ffauzann/simpleapi/internal/model/entity"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

func readConfig(conf *entity.Config) (err error) {
	b, err := ioutil.ReadFile("./internal/config/config.yaml")
	if err != nil {
		zap.S().Error(err)
		return
	}

	yaml.Unmarshal(b, conf)

	return
}
