package ex515

func Max(vals ...int) int {
	if len(vals) == 0 {
		panic("need least one number")
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func Min(vals ...int) int {
	if len(vals) == 0 {
		panic("need least one number")
	}
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func Join(sep string, elems ...string) string {
	var s string
	for _, v := range elems[0 : len(elems)-1] {
		s += v + sep
	}
	s += elems[len(elems)-1]
	return s
}
