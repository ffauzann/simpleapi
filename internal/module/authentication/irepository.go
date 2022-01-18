package authentication

import (
	"context"

	"github.com/ffauzann/simpleapi/internal/model/entity"
)

type Repository interface {
	// GetUserDetailByUsername gets a user detail including their merchant with given username parameter.
	GetUserDetailByUsername(ctx context.Context, username string) (userDetail entity.UserWithMerchant, err error)
}
