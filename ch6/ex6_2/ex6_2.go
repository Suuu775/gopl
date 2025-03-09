package ex62

type IntSet struct {
	words []uint64
}

func (s *IntSet) AddAll(values ...int) {
	for _, x := range values {
		s.Add(x)
	}
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}
