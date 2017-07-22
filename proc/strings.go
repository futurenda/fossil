package proc

import (
	"strings"
	"unicode"
)

func snakeToCamelCase(s string) string {
	output := ""
	for i, c := range s {
		if (i == 0 || string(s[i-1]) == "_") && string(s[i]) != "_" {
			output += strings.ToUpper(string(c))
		} else if !(string(s[i]) == "_") {
			output += strings.ToLower(string(c))
		}
	}
	return output
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func regularizeToSnakeCase(s string) string {
	regularized := ""
	for _, c := range s {
		if len(regularized) == 0 && unicode.IsDigit(c) {
			continue
		}
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			if isUpper(c) && !unicode.IsDigit(c) {
				if len(regularized) != 0 {
					regularized += "_"
				}
				regularized += strings.ToLower(string(c))
			} else {
				regularized += strings.ToLower(string(c))
			}
		} else {
			if len(regularized) != 0 || string(c) == "_" {
				regularized += "_"
			}
		}
	}
	return regularized
}
