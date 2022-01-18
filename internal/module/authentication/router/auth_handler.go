package router

import (
	mreq "github.com/ffauzann/simpleapi/internal/model/request"
	mres "github.com/ffauzann/simpleapi/internal/model/response"
	"github.com/ffauzann/simpleapi/internal/util/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Login godoc
// @Summary login
// @Tags authentication
// @Accept json
// @Produce json
// @Param payload body request.Login true "payload"
// @Success 200 {object} response.ExampleLogin "OK"
// @Router /login [POST]
func (h *Handler) Login(c echo.Context) (err error) {
	ctx := c.Request().Context()
	req := mreq.Login{}
	res := mres.Login{}

	err = c.Bind(&req)
	if err != nil {
		zap.S().Error(err)
		return response.Error(c, err)
	}

	res, err = h.Service.Login(ctx, &req)
	if err != nil {
		zap.S().Error(err)
		return response.Error(c, err)
	}

	return response.Success(c, res)
}
