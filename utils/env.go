package utils

import (
	"fmt"
	"regexp"
	"strings"
)

var blacklistedVars = []string{
	"PATH",
}

// CleanEnvArray remove blacklisted variables and format the result
func CleanEnvArray(envArray []string) []string {
	var resultArray []string
	for _, envVar := range envArray {
		varParts := strings.Split(envVar, "=")
		if isListed, _ := isBlacklisted(varParts[0]); !isListed {
			resultArray = append(resultArray, cleanValue(varParts[0], varParts[1]))
		}
	}
	return resultArray
}

// isBlacklisted return true if the item is blacklisted
func isBlacklisted(item string) (exists bool, index int) {
	exists = false
	index = -1
	for i, val := range blacklistedVars {
		if val == item {
			index = i
			exists = true
			return
		}
	}
	return
}

// cleanValue remove tab character, carrier return and 2 space tab from the value.
// Also removes begin and end double quotes and scape double quotes from the value.
func cleanValue(key string, value string) string {
	replacer := regexp.MustCompile("\n|\t|(  )|^\"|\"$")

	value = replacer.ReplaceAllString(value, "")
	value = strings.ReplaceAll(value, "\"", "\\\"")

	return fmt.Sprintf("%s=\"%s\"", key, value)
}
