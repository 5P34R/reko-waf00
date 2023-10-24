package util

import (
	"bytes"
	"strings"
)

func ConvertBufferToSlice(buffer *bytes.Buffer) []string {
	// Convert *bytes.Buffer to []string
	outputLines := strings.Split(buffer.String(), "\n")

	// Remove leading and trailing whitespace from each line
	for i, line := range outputLines {
		outputLines[i] = strings.TrimSpace(line)
	}

	return outputLines
}
