package main

import (
	"regexp"
	"strconv"
)

func Validate(cardNumber string) bool {
	parsedNumber := regexp.MustCompile(`[[:space:]|\p{Z}]`).ReplaceAllString(cardNumber, "")

	if !isValidNumber(parsedNumber) {
		return false
	}

	var (
		sum           int
		isSecondDigit bool
	)

	for i := len(parsedNumber) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(parsedNumber[i]))
		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		isSecondDigit = !isSecondDigit
	}

	return sum%10 == 0
}

func isValidNumber(cardNumber string) bool {
	return len(cardNumber) > 1 && regexp.MustCompile(`^[[:digit:]]*$`).MatchString(cardNumber)
}
