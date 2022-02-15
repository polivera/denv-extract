package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	envFileName string = ".env"
)

func WriteToEnvFile(envVarsList []string) (string, error) {
	var (
		filePath string
		err      error
		f        *os.File
	)

	// Clean values
	for ind, envVar := range envVarsList {
		varParts := strings.Split(envVar, "=")
		envVarsList[ind] = fmt.Sprintf("%s=\"%s\"", varParts[0], cleanEnvValue(varParts[1]))
	}
	// Get execution path
	if filePath, err = os.Getwd(); err != nil {
		return "", err
	}
	filePath = fmt.Sprintf("%s/%s", filePath, envFileName)
	// Create file
	if f, err = os.Create(filePath); err != nil {
		return "", err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			fmt.Printf("cant close file. Error: %s\n", err.Error())
		}
	}(f)
	// Write data
	if _, err = f.WriteString(strings.Join(envVarsList, "\n")); err != nil {
		return "", err
	}
	return filePath, nil
}

func ReadFromEnvFile(path string) ([]string, error) {
	var (
		content []byte
		err     error
		result  []string
	)

	if content, err = os.ReadFile(path); err != nil {
		fmt.Println("Cannot read file") //Should this be panic?
		return nil, err
	}

	for _, line := range strings.Split(string(content), "\n") {
		result = append(result, line)
	}

	return result, nil
}

// cleanEnvValue remove spaces, tabs and carrier returns from env var values
func cleanEnvValue(envValue string) string {
	replacer := regexp.MustCompile("\n|\t|(    )")

	envValue = strings.ReplaceAll(envValue, "\"", "\\\"")
	envValue = replacer.ReplaceAllString(envValue, "")
	return envValue
}
