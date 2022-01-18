package service

import (
	"context"

	"github.com/ffauzann/simpleapi/internal/constant"
	"github.com/ffauzann/simpleapi/internal/helper"
	"github.com/ffauzann/simpleapi/internal/model/entity"
	"github.com/ffauzann/simpleapi/internal/model/request"
	mres "github.com/ffauzann/simpleapi/internal/model/response"
	"github.com/ffauzann/simpleapi/internal/module/authentication"
	"go.uber.org/zap"
)

type Service struct {
	Repository authentication.Repository
	Config     entity.ConfigApp
}

func Init(r authentication.Repository, conf entity.ConfigApp) authentication.Service {
	return &Service{
		Repository: r,
		Config:     conf,
	}
}

func (s *Service) Login(ctx context.Context, req *request.Login) (res mres.Login, err error) {
	userDetail, err := s.Repository.GetUserDetailByUsername(ctx, req.Username)
	if err != nil {
		zap.S().Error(err)
		return
	}

	if !helper.IsPasswordMatch(req.Password, userDetail.User.Password) {
		err = constant.ErrInvalidCredentials
		return
	}

	res.Token, err = helper.GenerateToken(&userDetail, s.Config.JWT.Secret)
	if err != nil {
		zap.S().Error(err)
		return
	}

	return
}
