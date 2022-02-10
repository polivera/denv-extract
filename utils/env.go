package utils

import "strings"

var blacklistedVars = []string{
	"PATH",
}

func CleanEnvArray(envArray []string) []string {
	var resultArray []string
	for _, envVar := range envArray {
		varParts := strings.Split(envVar, "=")
		if isListed, _ := isBlacklisted(varParts[0]); !isListed {
			resultArray = append(resultArray, envVar)
		}
	}

	return resultArray
}

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
