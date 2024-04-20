package timeops

import (
	"time"
)

// IsValidDate checks if the given month and year are in the future/current
func IsFutureDate(month, year int) bool {
	now := time.Now()
	currentMonth, currentYear := int(now.Month()), now.Year()

	if year < currentYear || (year == currentYear && month < currentMonth) {
		return false
	}

	return true
}

// IsYearInRange checks if the given year is before certain years in the future
func IsYearInRange(year, maxValidYears int) bool {
	currentYear := time.Now().Year()
	return year >= currentYear && year <= currentYear+maxValidYears
}
