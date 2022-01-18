package transaction

import (
	"context"

	"github.com/ffauzann/simpleapi/internal/model/entity"
	"github.com/ffauzann/simpleapi/internal/model/request"
)

type Repository interface {
	// CountOutletByMerchantID counts total outlet(s) with given merchantID parameter.
	CountOutletByMerchantID(ctx context.Context, merchantID int) int

	// CalcMerchantGross calculates every merhant's gross profit daily and counts total data for pagination purpose.
	CalcMerchantGross(ctx context.Context, req *request.MerchantGross) (data []entity.MerchantGross, total int, err error)

	// CalcOutletGross calculates every outlet's gross profit daily and counts total data for pagination purpose.
	CalcOutletGross(ctx context.Context, req *request.OutletGross) (data []entity.OutletGross, total int, err error)
}
