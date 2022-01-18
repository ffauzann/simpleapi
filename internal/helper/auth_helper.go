package helper

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/ffauzann/simpleapi/internal/model/entity"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

// IsPasswordMatch determines whether given user's password
// is match with hashed password stored in database or not.
func IsPasswordMatch(plainPassword, hashedPassword string) bool {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(plainPassword)))
	return hash == hashedPassword
}

// GenerateToken generates signed string for further use.
func GenerateToken(userDetail *entity.UserWithMerchant, key string) (ss string, err error) {
	t := time.Now().Unix()
	claims := entity.JWTClaims{
		ID:       int(userDetail.User.ID),
		Name:     userDetail.User.Name,
		UserName: userDetail.User.UserName,
		Merchant: entity.MerchantClaims{
			ID:   int(userDetail.Merchant.ID),
			Name: userDetail.Merchant.Name,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t + 900, // 15 mins
			Issuer:    "simpleapi",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err = token.SignedString([]byte(key))
	if err != nil {
		zap.S().Error(err)
		return
	}

	return
}
