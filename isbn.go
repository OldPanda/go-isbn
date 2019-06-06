package isbn

import (
	"fmt"
	"strconv"
	"strings"
)

// Calculate check digit for ISBN-13
func calCheckDigitIsbn13(isbn13 string) (string, error) {
	multipliers := []int{1, 3}
	sum := 0
	for idx, char := range isbn13[:12] {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return "", fmt.Errorf("Failed to convert char to int: %s", string(char))
		}

		if idx%2 == 0 {
			sum += digit * multipliers[0]
		} else {
			sum += digit * multipliers[1]
		}
	}

	checkDigit := (10 - sum%10) % 10
	return strconv.Itoa(checkDigit), nil
}

// Calculate check digit for ISBN-10
func calCheckDigitIsbn10(isbn10 string) (string, error) {
	sum := 0
	for idx, char := range isbn10[:9] {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return "", fmt.Errorf("Failed to convert char to int: %s", string(char))
		}

		sum += digit * (10 - idx)
	}

	checkDigit := (11 - sum%11) % 11

	if checkDigit == 10 {
		return "X", nil
	}
	return strconv.Itoa(checkDigit), nil
}

// Validate whether given ISBN is valid or not. Support
// both ISBN-10 and ISBN-13
func Validate(isbn string) bool {
	if len(isbn) != 10 && len(isbn) != 13 {
		return false
	}

	var checkDigitStr string
	var err error
	switch len(isbn) {
	case 13:
		{
			if !strings.HasPrefix(isbn, "978") && !strings.HasPrefix(isbn, "979") {
				return false
			}

			checkDigitStr, err = calCheckDigitIsbn13(isbn)
			if err != nil {
				return false
			}
		}
	case 10:
		{
			checkDigitStr, err = calCheckDigitIsbn10(isbn)
			if err != nil {
				return false
			}
		}
	}
	return checkDigitStr == string(isbn[len(isbn)-1])
}

// ConvertToIsbn13 converts ISBN-10 to ISBN-13
func ConvertToIsbn13(isbn10 string) (string, error) {
	if len(isbn10) != 10 {
		return "", fmt.Errorf("ISBN with length 10 is required. Given: %s", isbn10)
	}

	if !Validate(isbn10) {
		return "", fmt.Errorf("Not a valid ISBN-10: %s", isbn10)
	}

	first12Digits := "978" + string(isbn10[:9])
	checkDigitStr, err := calCheckDigitIsbn13(first12Digits)
	if err != nil {
		return "", fmt.Errorf("Failed to calculate check digit. Error: %v", err)
	}
	return first12Digits + checkDigitStr, nil
}

// ConvertToIsbn10 converts ISBN-13 to ISBN-10, if convertible
func ConvertToIsbn10(isbn13 string) (string, error) {
	if len(isbn13) != 13 {
		return "", fmt.Errorf("ISBN with length 13 is required. Given: %s", isbn13)
	}

	if !strings.HasPrefix(isbn13, "978") {
		return "", fmt.Errorf("Given ISBN-13 is not convertible to ISBN10: %s", isbn13)
	}

	if !Validate(isbn13) {
		return "", fmt.Errorf("Not a valid ISBN-13: %s", isbn13)
	}

	checkDigitStr, err := calCheckDigitIsbn10(string(isbn13[3:12]))
	if err != nil {
		return "", fmt.Errorf("Failed to calculate check digit. Error: %v", err)
	}

	return string(isbn13[3:12]) + checkDigitStr, nil
}
