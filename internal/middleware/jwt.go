package middleware

import (
	"strings"

	"github.com/ffauzann/simpleapi/internal/constant"
	"github.com/ffauzann/simpleapi/internal/model/entity"
	"github.com/ffauzann/simpleapi/internal/util/response"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Middlewares verifies user's credentials.
func JWT(secret string) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearer := c.Request().Header.Get("Authorization")
			c.Set("bearer", bearer)

			token, err := extractBearer(bearer)
			if err != nil {
				return response.Error(c, err)
			}

			claims, err := verifyToken(c, token, secret)
			if err != nil {
				err = constant.ErrUnauthorized
				return response.Error(c, err)
			}

			user := entity.JWTClaims{
				ID:       int(claims["id"].(float64)),
				Name:     claims["name"].(string),
				UserName: claims["user_name"].(string),
				Merchant: entity.MerchantClaims{
					ID:   int(claims["merchant.id"].(float64)),
					Name: claims["merchant.name"].(string),
				},
			}

			c.Set(constant.JWTKeyUser, user)

			return hf(c)
		}
	}
}

// extractBearer extracts token from bearer.
func extractBearer(bearer string) (token string, err error) {
	if bearer == "" {
		err = constant.ErrUnauthorized
		return
	}

	splittedBearer := strings.Split(bearer, " ")
	if len(splittedBearer) != 2 {
		err = constant.ErrUnauthorized
		return
	}

	token = splittedBearer[1]

	return
}

// verifyToken verifies token and parse claims.
func verifyToken(c echo.Context, token, secret string) (claims jwt.MapClaims, err error) {
	jwtToken, err := jwt.Parse(
		token,
		func(t *jwt.Token) (i interface{}, err error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return
			}
			i = []byte(secret)
			return
		},
	)
	if err != nil {
		zap.S().Error(err)
		err = constant.ErrUnauthorized
		return
	}

	claims = jwtToken.Claims.(jwt.MapClaims)
	return
}
