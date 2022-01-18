package entity

import (
	"github.com/golang-jwt/jwt"
)

type (
	JWTClaims struct {
		ID       int            `json:"id"`
		Name     string         `json:"name"`
		UserName string         `json:"user_name"`
		Merchant MerchantClaims `json:"merchant"`
		jwt.StandardClaims
	}

	MerchantClaims struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
