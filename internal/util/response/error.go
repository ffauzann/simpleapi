package response

import (
	"fmt"
	"net/http"

	"github.com/ffauzann/simpleapi/internal/constant"
	mresponse "github.com/ffauzann/simpleapi/internal/model/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Error determines which message should be sent.
func Error(c echo.Context, err error) error {
	// Check whether the error comes from validator
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return errorValidation(c, err)
	}

	// Check whether the error comes from validator
	if _, ok := err.(validator.ValidationErrors); ok {
		return errorValidation(c, err)
	}

	// Check whether the error comes from echo binder
	if _, ok := err.(*echo.HTTPError); ok {
		return errorBind(c, err)
	}

	// Default error code value
	var code int = 500

	// Determine which known error occurred
	for e := range constant.MapError {
		if err == e {
			code = constant.MapError[e]
			break
		}
	}

	// Masking error message for unknown error
	if code == 500 {
		// Debugging purpose
		zap.S().Error(err)
		err = constant.ErrInternalServerError
	}

	// Build response
	res := mresponse.Response{
		Meta: mresponse.Meta{
			StatusCode: code,
			Message:    err.Error(),
		},
		Data: nil,
	}

	return c.JSONPretty(code, res, "    ")
}

// errorBind sends 422 status code due invalid data type.
func errorBind(c echo.Context, err error) error {
	code := http.StatusUnprocessableEntity
	res := mresponse.Response{
		Data: nil,
		Meta: mresponse.Meta{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed to parse data: invalid data type",
		},
	}

	return c.JSONPretty(code, res, "    ")

}

// errorBind sends 400 status code due to some parameter doesn't met requirement.
func errorValidation(c echo.Context, err error) error {
	res := mresponse.Response{
		Data: nil,
		Meta: mresponse.Meta{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		},
	}

	var msg string

	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		switch err.Tag() {
		case "required":
			msg = fmt.Sprintf("%s is required", field)
		case "number":
			msg = fmt.Sprintf("%s must be numbers only", field)
		case "gte":
			msg = fmt.Sprintf("%s value must be greater than %s", field, err.Param())
		case "lte":
			msg = fmt.Sprintf("%s value must be lower than %s", field, err.Param())
		case "datetime":
			msg = fmt.Sprintf("invalid date format of %s: %s", field, err.Param())
		}
	}
	res.Meta.Message = msg

	return c.JSONPretty(http.StatusUnprocessableEntity, res, "    ")
}
