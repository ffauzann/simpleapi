package service

import (
	"github.com/ffauzann/simpleapi/internal/model/entity"
	"github.com/ffauzann/simpleapi/internal/module/transaction"
)

type Service struct {
	Repository transaction.Repository
	Config     entity.ConfigApp
}

func Init(r transaction.Repository, conf entity.ConfigApp) transaction.Service {
	return &Service{
		Repository: r,
		Config:     conf,
	}
}
