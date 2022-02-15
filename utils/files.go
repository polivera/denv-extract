package utils

import (
	"fmt"
	"os"
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
		return nil, err
	}

	for _, line := range strings.Split(string(content), "\n") {
		if line != "" {
			result = append(result, line)
		}
	}

	return result, nil
}
