package router

import (
	"github.com/ffauzann/simpleapi/internal/model/entity"
	"github.com/ffauzann/simpleapi/internal/module/authentication"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service authentication.Service
	Config  entity.ConfigApp
}

func Init(g *echo.Group, s authentication.Service, c entity.ConfigApp) {
	handler := Handler{
		Service: s,
		Config:  c,
	}

	g.POST("/login", handler.Login)
}
