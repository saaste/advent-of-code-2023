package utils

import (
	"fmt"
	"os"
)

func ReadContent(inputFile string, testFile string, test bool) (string, error) {
	if test {
		bytes, err := os.ReadFile(testFile)
		if err != nil {
			return "", fmt.Errorf("failed to read the test file %s", testFile)
		}
		return string(bytes), nil
	}

	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read the input file %s", inputFile)
	}
	return string(bytes), nil
}
