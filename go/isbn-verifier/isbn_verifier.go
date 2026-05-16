package isbnverifier

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {

	noDash := strings.ReplaceAll(isbn, "-", "")
	if len(noDash) != 10 {
		return false
	}

	sum := 0

	for i := 0; i < 9; i++ {
		char := rune(noDash[i])
		if !unicode.IsDigit(char) {
			return false
		}
		sum += int(char-'0') * (10 - i)
	}

	lastChar := rune(noDash[9])
	if unicode.IsDigit(lastChar) {
		sum += int(lastChar - '0')
	} else if lastChar == 'X' || lastChar == 'x' {
		sum += 10
	} else {
		return false
	}

	return sum%11 == 0
}
