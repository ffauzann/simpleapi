package authentication

import (
	"context"

	"github.com/ffauzann/simpleapi/internal/model/request"
	"github.com/ffauzann/simpleapi/internal/model/response"
)

type Service interface {
	Login(ctx context.Context, req *request.Login) (res response.Login, err error)
}
