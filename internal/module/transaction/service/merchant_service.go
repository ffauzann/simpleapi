package service

import (
	"context"

	"github.com/ffauzann/simpleapi/internal/model/request"
	"github.com/ffauzann/simpleapi/internal/model/response"
	"go.uber.org/zap"
)

// MerchantGross nodoc. See: iservice.go.
func (s *Service) MerchantGross(ctx context.Context, req *request.MerchantGross) (res []response.MerchantGross, err error) {
	// Calculate merchant's gross profit daily
	data, count, err := s.Repository.CalcMerchantGross(ctx, req)
	if err != nil {
		zap.S().Error()
		return
	}

	// Build response
	for i := range data {
		res = append(res, response.MerchantGross{
			Date: data[i].Date,
			Merchant: response.Merchant{
				ID:   data[i].MerchantID,
				Name: data[i].MerchantName,
			},
			Gross: data[i].Gross,
		})
	}

	req.Total = count // Set total data

	return
}
