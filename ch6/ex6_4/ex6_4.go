package ex64

type IntSet struct {
	words []uint64
}

func (s *IntSet) Elems() []int {
	var elements []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elements = append(elements, 64*i+j)
			}
		}
	}
	return elements
}
