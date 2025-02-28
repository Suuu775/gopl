package ex44

func Rotate(s []int, k int) {
	n := len(s)
	k = k % n

	for start := 0; start < gcd(n, k); start++ {
		current := start
		prev := s[start]

		for {
			next := (current + k) % n
			temp := s[next]
			s[next] = prev
			prev = temp
			current = next

			if current == start {
				break
			}
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
