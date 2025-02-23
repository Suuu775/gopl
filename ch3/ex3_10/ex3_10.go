package ex310

import "bytes"

func Comma(s string) string {
	var buf bytes.Buffer
	n := len(s)

	for i := n; i > 0; i -= 3 {
		if i-3 > 0 {
			buf.WriteString(s[i-3 : i])
			buf.WriteString(",")
		} else {
			buf.WriteString(s[:i])
		}
	}

	return buf.String()
}
