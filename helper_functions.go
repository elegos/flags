package flags

import (
	"fmt"
	"regexp"
	"strings"
)

func isOption(arg string) bool {
	return regexp.MustCompile("^--?[a-zA-Z0-9]{1,}").MatchString(arg)
}

func isShortOption(arg string) bool {
	return regexp.MustCompile("^-[a-zA-Z0-9]").MatchString(arg)
}

func getOptionName(arg string) (string, error) {
	matches := regexp.MustCompile("^--?([a-zA-Z0-9]{1,})").FindStringSubmatch(arg)
	// 2 = entire string + matched string
	matchCheckNum := 2
	if len(matches) != matchCheckNum {
		return "", fmt.Errorf(`"%s" is not an option`, arg)
	}

	return matches[1], nil
}

func textSplit(source string, maxChars int) ([]string, error) {
	result := []string{}

	if maxChars < 1 {
		return result, fmt.Errorf("maxChars must be >= 1")
	}

	buffer := ""

	for _, word := range strings.Split(source, " ") {
		if len(buffer) >= maxChars {
			result = append(result, strings.Trim(buffer, " "))
			buffer = ""
		}

		buffer = strings.Join(append(strings.Split(buffer, " "), word), " ")

		continue
	}

	if buffer != "" {
		result = append(result, strings.Trim(buffer, " "))
	}

	return result, nil
}
