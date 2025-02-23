package main

import (
	"bytes"
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.
func Comma(s string) string {
	var buf bytes.Buffer

	// Handle negative sign if present
	var sign string
	if strings.HasPrefix(s, "-") {
		sign = "-"
		s = s[1:]
	} else if strings.HasPrefix(s, "+") {
		sign = "+"
		s = s[1:]
	}

	// Split into integer and fractional parts
	parts := strings.Split(s, ".")
	integerPart := parts[0]
	var fractionalPart string
	if len(parts) > 1 {
		fractionalPart = parts[1]
	}

	// Process integer part
	n := len(integerPart)
	for i := n; i > 0; i -= 3 {
		if i-3 > 0 {
			buf.WriteString(integerPart[i-3 : i])
			buf.WriteString(",")
		} else {
			buf.WriteString(integerPart[:i])
		}
	}

	// Add fractional part if present
	if fractionalPart != "" {
		buf.WriteString(".")
		buf.WriteString(fractionalPart)
	}

	// Prepend sign if present
	if sign != "" {
		buf.WriteString(sign)
	}

	return buf.String()
}
