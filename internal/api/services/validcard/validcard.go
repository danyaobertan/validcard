package validcard

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"unicode"

	"github.com/danyaobertan/validcard/internal/api/domain"
	"github.com/danyaobertan/validcard/pkg/errorops"
	"github.com/danyaobertan/validcard/pkg/luhn"
	"github.com/danyaobertan/validcard/pkg/timeops"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// IsValidCardInfo validates card number and expiration date
func (s *Service) IsValidCardInfo(card domain.Card) (bool, *errorops.Error) {
	// validate expiration date
	valid, err := isValidExpDate(card.ExpirationMonth, card.ExpirationYear)
	if !valid {
		return false, err
	}

	// validate card number
	valid, err = isValidCardNumber(card.CardNumber)
	if !valid {
		return false, err
	}

	return true, nil
}

// IsValidDate checks if the given month and year are in the future
func isValidExpDate(month, year int) (bool, *errorops.Error) {
	// validate month
	if month < 1 || month > 12 {
		return false, errorops.NewError(
			http.StatusBadRequest,
			"invalid month",
			"month must be between 1 and 12",
		)
	}

	// validate year
	if !timeops.IsYearInRange(year, domain.MaxYearsInFuture) {
		return false, errorops.NewError(
			http.StatusBadRequest,
			"invalid year",
			fmt.Sprintf("year must be between %d and %d", time.Now().Year(), time.Now().Year()+domain.MaxYearsInFuture),
		)
	}

	// validate expiration date
	if !timeops.IsFutureDate(month, year) {
		return false, errorops.NewError(
			http.StatusBadRequest,
			"invalid expiration date",
			"expiration date must be in the future",
		)
	}

	return true, nil
}

func isValidCardNumber(cardNumber string) (bool, *errorops.Error) {
	// validate card number length
	if len(cardNumber) < domain.CardNumberMinLength || len(cardNumber) > domain.CardNumberMaxLength {
		return false, errorops.NewError(
			http.StatusBadRequest,
			"invalid card number length",
			fmt.Sprintf("card number length must be between %d and %d, but current length is %d", domain.CardNumberMinLength, domain.CardNumberMaxLength, len(cardNumber)),
		)
	}

	// strip spaces, check for invalid characters and convert to digit slice
	digits, vtErr := validateTransform(cardNumber)
	if vtErr != nil {
		return false, errorops.NewError(
			http.StatusBadRequest,
			"invalid card number symbols",
			fmt.Sprintf("card number could contain only digits and spaces, but contains %s", vtErr.Error()),
		)
	}

	// validate card number using Luhn algorithm
	isValidLuhn := luhn.IsValid(digits)
	if !isValidLuhn {
		return false, errorops.NewError(
			http.StatusBadRequest,
			"invalid card number by Luhn algorithm",
			nil,
		)
	}

	return true, nil
}

// validateTransform validates card number for symbols
func validateTransform(number string) ([]int, error) {
	var digits []int

	for _, d := range number {
		if unicode.IsDigit(d) {
			digit, err := strconv.Atoi(string(d))
			if err != nil {
				return nil, fmt.Errorf("failed to convert digit to int: %s", string(d))
			}

			digits = append(digits, digit)
		} else if !unicode.IsSpace(d) {
			return nil, fmt.Errorf("invalid character: %s", string(d))
		}
	}

	return digits, nil
}
