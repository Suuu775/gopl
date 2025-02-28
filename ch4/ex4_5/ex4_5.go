package ex45

func Dedup(str []string) []string {
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			copy(str[i:], str[i+1:])
			str = str[:len(str)-1]
		}
	}
	return str
}
