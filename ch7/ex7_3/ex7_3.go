package ex73

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = root.add(v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func (t *tree) add(value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = t.left.add(value)
	} else {
		t.right = t.right.add(value)
	}
	return t
}

func (s *tree) String() string {
	if s == nil {
		return ""
	} else {
		return fmt.Sprintf("%s %s %s", fmt.Sprint(s.value), s.left.String(), s.right.String())
	}

}
