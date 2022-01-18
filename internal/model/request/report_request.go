package request

import "github.com/ffauzann/simpleapi/internal/model/entity"

type ReportPagination struct {
	StartDate string `query:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate   string `query:"end_date" validate:"required,datetime=2006-01-02"`
	entity.Pagination
	User entity.JWTClaims
}
