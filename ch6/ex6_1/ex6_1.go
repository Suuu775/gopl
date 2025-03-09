package ex61

// import (
// 	"math/bits"

// 	"github.com/Suuu775/gopl/ch6/intset/"
// )

// func (s *intset.IntSet) Len() int {
// 	count := 0
// 	for _, word := range s.words {
// 		count += int(bits.OnesCount64(word))
// 	}
// 	return count
// }

// func (s *intset.IntSet) Remove(x int) {
// 	word, bit := x/64, uint(x%64)
// 	if word < len(s.words) {
// 		s.words[word] &^= 1 << bit
// 	}
// }

// func (s *intset.IntSet) Clear() {
// 	s.words = s.words[:0]
// }

// func (s *intset.IntSet) Copy() *intset.IntSet {
// 	copySet := &intset.IntSet{
// 		words: make([]uint64, len(s.words)),
// 	}
// 	copy(copySet.words, s.words)
// 	return copySet
// }
