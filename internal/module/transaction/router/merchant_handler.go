package router

import (
	"github.com/ffauzann/simpleapi/internal/constant"
	"github.com/ffauzann/simpleapi/internal/model/entity"
	mreq "github.com/ffauzann/simpleapi/internal/model/request"
	"github.com/ffauzann/simpleapi/internal/util/request"
	"github.com/ffauzann/simpleapi/internal/util/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// MerchantGross counts merchant's gross income daily
// @Summary count merchant's gross income daily
// @Tags merchant
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer + token" default(Bearer {token})
// @Param start_date query string true "Format: yyyy-mm-dd" default(2021-11-01)
// @Param end_date query string true "Format: yyyy-mm-dd" default(2021-11-30)
// @Param page query int true "Page" default(1)
// @Param limit query int true "Limit" default(10)
// @Success 200 {object} response.ExampleMerchantGross "OK"
// @Router /merchant/gross [GET]
func (h *Handler) MerchantGross(c echo.Context) (err error) {
	ctx := c.Request().Context()
	cre := c.Get(constant.JWTKeyUser).(entity.JWTClaims)
	req := mreq.MerchantGross{
		ReportPagination: mreq.ReportPagination{
			User: cre,
		},
	}

	if ok := request.BindAndValidate(c, &req, h.Config.Validator); !ok {
		return
	}

	req.Pagination.Calc()

	res, err := h.Service.MerchantGross(ctx, &req)
	if err != nil {
		zap.S().Error(err)
		return response.Error(c, err)
	}

	return response.Success(c, res, req.Pagination)
}
