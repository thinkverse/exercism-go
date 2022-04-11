package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	if err != nil {
		panic(err)
	}

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`\<[(*=~\-)]*\>`)

	if err != nil {
		panic(err)
	}

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`"(.*password.*)"`)

	if err != nil {
		panic(err)
	}

	count := 0

	for _, line := range lines {
		count += len(re.FindStringSubmatch(line))
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line\d+`)

	if err != nil {
		panic(err)
	}

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, err := regexp.Compile(`(?:User\s+([A-z0-9]+))`)

	if err != nil {
		panic(err)
	}

	for index, line := range lines {
		if re.MatchString(line) {
			submatch := re.FindStringSubmatch(line)
			lines[index] = fmt.Sprintf("[USR] %s %s", submatch[1], line)
		}
	}

	return lines
}
