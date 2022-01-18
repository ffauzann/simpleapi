package request

import (
	"github.com/ffauzann/simpleapi/internal/util/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// BindAndValidate binds and validate request parameter.
func BindAndValidate(e echo.Context, i interface{}, v *validator.Validate) (ok bool) {
	err := e.Bind(i)
	if err != nil {
		zap.S().Error(err)
		response.Error(e, err)
		return false
	}

	err = v.Struct(i)
	if err != nil {
		zap.S().Error(err)
		response.Error(e, err)
		return false
	}

	return true
}
