package service

import (
	"context"

	"github.com/ffauzann/simpleapi/internal/constant"
	"github.com/ffauzann/simpleapi/internal/model/request"
	"github.com/ffauzann/simpleapi/internal/model/response"
	"go.uber.org/zap"
)

// OutletGross nodoc. See: iservice.go.
func (s *Service) OutletGross(ctx context.Context, req *request.OutletGross) (res []response.OutletGross, err error) {
	// Count total outlet that user/merchant has. if there is no outlet, throw an error.
	// Otherwise, if the merchant are more than 1, increase limit and offset.
	totalOutlet := s.Repository.CountOutletByUserID(ctx, req.User.ID)
	if totalOutlet == 0 {
		err = constant.ErrNoOutletFound
		return
	} else if totalOutlet > 1 {
		req.Pagination.Limit *= totalOutlet
		if req.Pagination.Page > 1 {
			req.Pagination.Offset *= totalOutlet
		}
	}

	// Calculate outlet's gross profit daily
	data, count, err := s.Repository.CalcOutletGross(ctx, req)
	if err != nil {
		zap.S().Error()
		return
	}

	// Build response data
	for i := range data {
		// Prepare outlet detail
		outlet := response.OutletWithGross{
			ID:    data[i].OutletID,
			Name:  data[i].OutletName,
			Gross: data[i].Gross,
		}

		// Append only the outlet
		if i != 0 {
			if data[i].Date == data[i-1].Date {
				res[len(res)-1].Merchant.Outlets = append(res[len(res)-1].Merchant.Outlets, outlet)
				continue
			}
		}

		// Prepare merchant detail
		merchant := response.MerchantWithOutlets{
			ID:      data[i].MerchantID,
			Name:    data[i].MerchantName,
			Outlets: []response.OutletWithGross{outlet},
		}

		// Append new date
		res = append(res, response.OutletGross{
			Date:     data[i].Date,
			Merchant: merchant,
		})
	}

	req.Pagination.Limit /= totalOutlet // Set limit back
	req.Total = count / totalOutlet     // Set total data

	return
}
