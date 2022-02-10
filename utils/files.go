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
	defer f.Close()

	// Write data
	if _, err = f.WriteString(strings.Join(envVarsList, "\n")); err != nil {
		return "", err
	}

	return filePath, nil
}
