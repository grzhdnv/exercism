// Package phonenumber implements the phone number exercise.
package phonenumber

import "errors"

// Number cleans up user-entered phone numbers so that they can be sent as SMS messages. The North American Numbering Plan (NANP) is a telephone numbering system used by many countries in North America, including the United States and Canada. NANP numbers are ten digits long. The first three digits are the area code, followed by a three-digit exchange code, and then a four-digit subscriber number. NANP numbers can be written with various punctuation and the country code of 1 is sometimes prepended. Your task is to clean up differently formatted telephone numbers by removing punctuation and the country code (if present).
func Number(phoneNumber string) (string, error) {

	digits := ""
	// Extract digits from the input string
	for _, r := range phoneNumber {
		if r >= '0' && r <= '9' {
			digits += string(r)
		}
	}

	// Validate the number of digits and the country code
	if len(digits) != 10 {
		if len(digits) == 11 && digits[0] == '1' {
			digits = digits[1:]
		} else {
			return "", errors.New("invalid phone number")
		}
	}

	// Validate area code and exchange code
	if digits[0] < '2' || digits[3] < '2' {
		return "", errors.New("invalid phone number")
	}

	return digits, nil
}

// AreaCode extracts the area code from a cleaned-up phone number.
func AreaCode(phoneNumber string) (string, error) {
	cleaned, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return cleaned[:3], err
}

func Format(phoneNumber string) (string, error) {
	cleaned, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + cleaned[:3] + ") " + cleaned[3:6] + "-" + cleaned[6:], nil
}
