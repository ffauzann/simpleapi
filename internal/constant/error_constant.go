package constant

import (
	"fmt"
	"net/http"
)

var (
	// 400
	ErrNoOutletFound = fmt.Errorf("no outlet found")

	// 401
	ErrUnauthorized       = fmt.Errorf("%s", http.StatusText(http.StatusUnauthorized))
	ErrInvalidCredentials = fmt.Errorf("invalid username/password")

	// 500
	ErrInternalServerError = fmt.Errorf("internal server error")

	MapError = map[error]int{
		// 400
		ErrNoOutletFound: http.StatusBadRequest,

		// 401
		ErrInvalidCredentials: http.StatusUnauthorized,
		ErrUnauthorized:       http.StatusUnauthorized,

		// 500
		ErrInternalServerError: http.StatusInternalServerError,
	}
)
