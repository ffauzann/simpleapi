package router

import (
	"github.com/ffauzann/simpleapi/internal/model/entity"
	"github.com/ffauzann/simpleapi/internal/module/transaction"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service transaction.Service
	Config  entity.ConfigApp
}

func Init(g *echo.Group, s transaction.Service, c entity.ConfigApp) {
	handler := Handler{
		Service: s,
		Config:  c,
	}

	g.GET("/merchant/gross", handler.MerchantGross)
	g.GET("/merchant/outlet/gross", handler.OutletGross)
}
