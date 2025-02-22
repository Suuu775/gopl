package ex23

// pc[i] is the population count of i.
var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	res := 0
	for i := 0; i < 7; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

func PopCountIter(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// find 1 one bit by one bit
func PopCountIterOnce(x uint64) int {
	res := 0
	for ; x != 0; x >>= 1 {
		res += int(x & 1)
	}
	return res
}

// ex2-5
func SparsePopCount(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}
