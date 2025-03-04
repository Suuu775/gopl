package ex516

func Join(sep string, elems ...string) string {
	var s string
	for _, v := range elems[0 : len(elems)-1] {
		s += v + sep
	}
	s += elems[len(elems)-1]
	return s
}
