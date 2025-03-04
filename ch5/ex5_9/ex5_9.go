package ex59

import "strings"

func Expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "foo", f("foo"))
}
