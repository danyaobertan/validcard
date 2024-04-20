package luhn

// IsValid checks if the number is valid per the Luhn formula with the final multiplication by 9.
func IsValid(digits []int) bool {
	if len(digits) <= 1 {
		return false
	}

	// Last digit is the check digit.
	checkDigit := digits[len(digits)-1]

	// Process all digits except the check digit.
	digits = digits[:len(digits)-1]

	var sum int

	digitsLen := len(digits)

	for i := digitsLen - 1; i >= 0; i -= 2 {
		// Double every second digit from the right
		digits[i] *= 2
		if digits[i] > 9 { //nolint:gomnd // If the result is greater than 9, subtract 9
			digits[i] -= 9
		}
	}

	for _, digit := range digits {
		sum += digit
	}

	// Multiply the sum by 9 and take the last digit.
	result := (sum * 9) % 10 //nolint:gomnd // The last digit of the result is the check digit.

	// Check if the last digit of the result matches the check digit.
	return result == checkDigit
}
