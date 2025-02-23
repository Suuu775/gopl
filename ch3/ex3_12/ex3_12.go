package ex312

func IsSameString(s, s1 string) bool {

	if len(s) != len(s1) {
		return false
	}
	var map1 = make(map[rune]uint)
	var map2 = make(map[rune]uint)

	for _, v := range s {
		map1[v] += 1
	}

	for _, v := range s1 {
		map2[v] += 1
	}

	for k, v := range map1 {
		if val, ok := map2[k]; ok {
			if v == val {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
