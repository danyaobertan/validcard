package validcard

import (
	"net/http"

	"github.com/danyaobertan/validcard/pkg/errorops"
)

func (r RequestBody) validate() *errorops.Error {
	if r.CardNumber == "" {
		return errorops.NewError(
			http.StatusBadRequest,
			"card number is empty",
			nil,
		)
	}

	if r.ExpirationMonth == 0 {
		return errorops.NewError(
			http.StatusBadRequest,
			"expiration month is empty",
			nil,
		)
	}

	if r.ExpirationYear == 0 {
		return errorops.NewError(
			http.StatusBadRequest,
			"expiration year is empty",
			nil,
		)
	}

	return nil
}
