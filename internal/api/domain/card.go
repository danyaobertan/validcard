package domain

const (
	// CardNumberMinLength is the minimum length of a card number currently supported
	CardNumberMinLength = 12
	// CardNumberMaxLength is the maximum length of a card number currently supported
	CardNumberMaxLength = 19
	// MaxYearsInFuture is the maximum number of years in the future that a card expiration date can be
	MaxYearsInFuture = 20
)

type Card struct {
	CardNumber      string
	ExpirationMonth int
	ExpirationYear  int
}
