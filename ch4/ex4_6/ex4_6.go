package ex46

import "unicode"

func DedupSpace(s []byte) []byte {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] && unicode.IsSpace(rune(s[i])) {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		}
	}
	return s
}
