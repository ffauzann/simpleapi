package repository

import (
	"context"
	"fmt"

	"github.com/ffauzann/simpleapi/internal/model/entity"
	"github.com/ffauzann/simpleapi/internal/model/request"
	"go.uber.org/zap"
)

// CalcMerchantGross nodoc. See: irepository.go.
func (r *Repository) CalcMerchantGross(ctx context.Context, req *request.MerchantGross) (data []entity.MerchantGross, total int, err error) {
	baseQuery := fmt.Sprintf(
		`WITH RECURSIVE DateRange AS (
			SELECT '%s' date UNION SELECT date + INTERVAL 1 DAY
			FROM DateRange
			WHERE date < '%s'
		)
		SELECT dr.date, m.merchant_id, m.merchant_name, COALESCE(t.gross, 0) gross
		FROM DateRange dr
		LEFT JOIN (
			SELECT id merchant_id, merchant_name
			FROM Merchants
			WHERE user_id = %d
		) AS m ON merchant_id IS NOT NULL
		LEFT JOIN (
			SELECT merchant_id, SUM(bill_total) gross, LEFT(created_at, 10) date
			FROM Transactions
			GROUP BY merchant_id, LEFT(created_at, 10)
		) AS t ON t.merchant_id = m.merchant_id AND t.date = dr.date
		ORDER BY date, merchant_id`,
		req.StartDate,
		req.EndDate,
		req.User.ID,
	)
	limitOffset := fmt.Sprintf("LIMIT %d OFFSET %d", req.Pagination.Limit, req.Pagination.Offset)

	var count int64
	countQuery := fmt.Sprintf("SELECT count(date) FROM (%s) AS Source", baseQuery)
	r.DB.Raw(countQuery).Scan(&count)
	total = int(count)

	fullQuery := fmt.Sprintf("%s %s", baseQuery, limitOffset)
	rows, err := r.DB.Raw(fullQuery).Rows()
	if err != nil {
		zap.S().Error(rows)
		return
	}

	mg := entity.MerchantGross{}
	for rows.Next() {
		r.DB.ScanRows(rows, &mg)
		data = append(data, mg)
	}

	return
}
