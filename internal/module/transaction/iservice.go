package transaction

import (
	"context"

	"github.com/ffauzann/simpleapi/internal/model/request"
	"github.com/ffauzann/simpleapi/internal/model/response"
)

type Service interface {
	// MerchantGross calculates daily gross profit of a merchant owned by logged in user.
	// It counts 0 if the day within the range has no any transactions.
	MerchantGross(ctx context.Context, req *request.MerchantGross) (res []response.MerchantGross, err error)

	// OutletGross calculates daily gross profit of outlets owned by logged in user.
	// It counts 0 if the day within the range has no any transactions.
	OutletGross(ctx context.Context, req *request.OutletGross) (res []response.OutletGross, err error)
}
