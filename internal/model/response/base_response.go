package response

import "github.com/ffauzann/simpleapi/internal/model/entity"

type Meta struct {
	StatusCode int    `json:"status_code" example:"200"`
	Message    string `json:"message" example:"OK"`
}

type MetaPagination struct {
	Meta
	Pagination entity.Pagination `json:"pagination"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ResponsePagination struct {
	Meta MetaPagination `json:"meta"`
	Data interface{}    `json:"data"`
}
