package parser

import (
	"fmt"
	"strings"
)

func ParseFloatArray(s string) ([]float64, error) {
	// Remove curly braces
	s = strings.Trim(s, "{}")
	if s == "" {
		return nil, nil
	}

	// Split by comma
	parts := strings.Split(s, ",")

	// Convert to float64
	result := make([]float64, len(parts))
	for i, part := range parts {
		var value float64
		if _, err := fmt.Sscanf(part, "%f", &value); err != nil {
			return nil, err
		}
		result[i] = value
	}

	return result, nil
}
