package luhn

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
	"testing"
	"unicode"
)

type CreditCard struct {
	IssuingNetwork string `json:"IssuingNetwork"`
	CardNumber     int64  `json:"CardNumber"`
}

type CreditCardEntry struct {
	CreditCard CreditCard `json:"CreditCard"`
}

func LoadTestCasesFromFile(filePath string) ([]CreditCardEntry, error) {
	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into the Go structs
	var cards []CreditCardEntry
	err = json.Unmarshal(data, &cards)

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func Transform(number string) ([]int, error) {
	var digits []int

	for _, d := range number {
		if unicode.IsDigit(d) {
			digit, _ := strconv.Atoi(string(d))
			digits = append(digits, digit)
		} else if !unicode.IsSpace(d) {
			return nil, errors.New("invalid character in card number")
		}
	}

	return digits, nil
}

func TestValid(t *testing.T) {
	filePath := "lunn_test_cases.json"
	// Load credit card data from the specified file
	testCases, err := LoadTestCasesFromFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load test cases from file: %v", err)
	}

	for _, test := range testCases {
		digits, err := Transform(strconv.FormatInt(test.CreditCard.CardNumber, 10))

		if err != nil {
			t.Fatalf("Transform(%v): %s", test.CreditCard.CardNumber, err)
		}

		if ok := IsValid(digits); ok != true {
			t.Fatalf("Valid(%v): %s\n\t Expected: %t\n\t", test.CreditCard.CardNumber, test.CreditCard.IssuingNetwork, ok)
		}
	}
}

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValid([]int{4, 5, 3, 2, 0, 1, 5, 1, 1, 2, 8, 3, 0, 3, 6})
	}
}
