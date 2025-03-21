package ex63

import (
	"bytes"
	"fmt"
	"math/bits"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	if len(t.words) < len(s.words) {
		s.words = s.words[:len(t.words)]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	result := &IntSet{}
	maxLen := len(s.words)
	if len(t.words) > maxLen {
		maxLen = len(t.words)
	}
	for i := 0; i < maxLen; i++ {
		var word uint64
		if i < len(s.words) {
			word = s.words[i]
		}
		if i < len(t.words) {
			word ^= t.words[i]
		}
		result.words = append(result.words, word)
	}
	return result
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += int(bits.OnesCount64(word))
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

func (s *IntSet) Copy() *IntSet {
	copySet := &IntSet{
		words: make([]uint64, len(s.words)),
	}
	copy(copySet.words, s.words)
	return copySet
}

func (s *IntSet) AddAll(values ...int) {
	for _, x := range values {
		s.Add(x)
	}
}
