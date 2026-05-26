package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	matches := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, l := range lines {
		if re.MatchString(l) {
			matches++
		}
	}
	return matches
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)

	tagged := make([]string, len(lines))
	for i, line := range lines {
		if m := re.FindStringSubmatch(line); m != nil {
			tagged[i] = fmt.Sprintf("[USR] %s %s", m[1], line)
		} else {
			tagged[i] = line
		}
	}
	return tagged
}
