package response

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/ffauzann/simpleapi/internal/model/entity"
	mresponse "github.com/ffauzann/simpleapi/internal/model/response"

	"github.com/labstack/echo/v4"
)

func Success(c echo.Context, data interface{}, pagination ...entity.Pagination) (err error) {
	if len(pagination) == 0 {
		return simpleResponse(c, data)
	}
	return paginateResponse(c, data, pagination[0])
}

func simpleResponse(c echo.Context, data interface{}) (err error) {
	res := mresponse.Response{
		Meta: mresponse.Meta{
			StatusCode: http.StatusOK,
			Message:    http.StatusText(http.StatusOK),
		},
		Data: data,
	}

	return c.JSON(http.StatusOK, res)
}

func paginateResponse(c echo.Context, data interface{}, pagination entity.Pagination) (err error) {
	setPagination(c, &pagination)
	res := mresponse.ResponsePagination{
		Meta: mresponse.MetaPagination{
			Meta: mresponse.Meta{
				StatusCode: http.StatusOK,
				Message:    http.StatusText(http.StatusOK),
			},
			Pagination: pagination,
		},
		Data: data,
	}

	return c.JSON(http.StatusOK, res)
}

func setPagination(c echo.Context, p *entity.Pagination) {
	p.From = (p.Page-1)*10 + 1
	p.To = int(int(p.Page) * int(p.Limit))
	p.LastPage = int(math.Ceil(float64(p.Total) / float64(p.Limit)))

	p.CurrentPageURL = fmt.Sprintf("%s%v", c.Request().Host, c.Request().URL.String())

	if p.Page > 1 {
		c.QueryParams().Set("page", strconv.Itoa(p.Page-1))
		p.PrevPageURL = fmt.Sprintf("%s%s?%s", c.Request().Host, c.Path(), c.QueryParams().Encode())
	}

	if p.Page < p.LastPage {
		c.QueryParams().Set("page", strconv.Itoa(p.Page+1))
		p.NextPageURL = fmt.Sprintf("%s%s?%s", c.Request().Host, c.Path(), c.QueryParams().Encode())
	}
}
