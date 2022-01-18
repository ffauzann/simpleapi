package repository

import (
	"context"
	"fmt"

	"github.com/ffauzann/simpleapi/internal/model/entity"
	"github.com/ffauzann/simpleapi/internal/model/request"
	"go.uber.org/zap"
)

// CountOutletByUserID nodoc. See: irepository.go.
func (r *Repository) CountOutletByUserID(ctx context.Context, userID int) int {
	var count int64
	query := fmt.Sprintf(
		`SELECT o.id 
		FROM Outlets o
		RIGHT JOIN (
			SELECT id
			FROM Merchants
			WHERE user_id = %d
		) as m ON m.id = o.merchant_id`,
		userID,
	)
	r.DB.Raw(query).Count(&count)
	return int(count)
}

// CalcOutletGross nodoc. See: irepository.go.
func (r *Repository) CalcOutletGross(ctx context.Context, req *request.OutletGross) (data []entity.OutletGross, total int, err error) {
	baseQuery := fmt.Sprintf(
		`WITH RECURSIVE DateRange AS (
			SELECT '%s' date UNION SELECT date + INTERVAL 1 DAY
			FROM DateRange
			WHERE date < '%s'
		)
		SELECT dr.date, o.merchant_id, o.merchant_name, o.outlet_id, o.outlet_name, COALESCE(t.gross, 0) gross
		FROM DateRange dr
		LEFT JOIN (
			SELECT m.id merchant_id, m.merchant_name, o.id outlet_id, outlet_name
			FROM Outlets o
			LEFT JOIN (
				SELECT id, merchant_name
				FROM Merchants
				WHERE user_id = %d
			) AS m ON m.id = o.merchant_id
		) AS o ON merchant_id IS NOT NULL
		LEFT JOIN (
			SELECT outlet_id, SUM(bill_total) gross, LEFT(created_at, 10) date
			FROM Transactions
			WHERE merchant_id IN (SELECT id FROM Merchants WHERE user_id = %d)
			GROUP BY outlet_id, LEFT(created_at, 10)
		) AS t ON t.outlet_id = o.outlet_id AND t.date = dr.date
		ORDER BY date, outlet_id`,
		req.StartDate,
		req.EndDate,
		req.User.ID,
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

	og := entity.OutletGross{}
	for rows.Next() {
		r.DB.ScanRows(rows, &og)
		data = append(data, og)
	}

	return
}
