package luhn

import (
	"unicode"
)

func Valid(id string) bool {

	sum := 0
	count := 0

	for i := len(id) - 1; i >= 0; i-- {

		char := rune(id[i])

		if !unicode.IsDigit(char) && !unicode.IsSpace(char) {
			return false
		}

		if unicode.IsDigit(char) {
			digit := int(char - '0')
			if count%2 == 1 {
				digit = digit * 2
				if digit > 9 {
					digit -= 9
				}
			}
			sum += digit
			count++
		}
	}
	return count > 1 && sum%10 == 0
}
