package ex23_test

import (
	"testing"

	ex23 "github.com/Suuu775/gopl/ch2/ex2_3"
)

func BenchmarkPopCount(b *testing.B) {
	b.Run("PopCount use single express", func(b *testing.B) {
		ex23.PopCount(18446744073709551615)
	})
	b.Run("PopCount use iter", func(b *testing.B) {
		ex23.PopCountIter(18446744073709551615)
	})
	// ex2-4
	b.Run("PopCount use iter one bit by one bit", func(b *testing.B) {
		ex23.PopCountIterOnce(18446744073709551615)
	})
	b.Run("PopCount use n &= n-1", func(b *testing.B) {
		ex23.SparsePopCount(18446744073709551615)
	})
}

// Running tool: D:\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkPopCount$ github.com/Suuu775/gopl/ch2/ex2_3

// goos: windows
// goarch: amd64
// pkg: github.com/Suuu775/gopl/ch2/ex2_3
// cpu: 12th Gen Intel(R) Core(TM) i5-12500H
// === RUN   BenchmarkPopCount
// BenchmarkPopCount
// === RUN   BenchmarkPopCount/PopCount_use_single_express
// BenchmarkPopCount/PopCount_use_single_express
// BenchmarkPopCount/PopCount_use_single_express-16
// 1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
// === RUN   BenchmarkPopCount/PopCount_use_iter
// BenchmarkPopCount/PopCount_use_iter
// BenchmarkPopCount/PopCount_use_iter-16
// 1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
// === RUN   BenchmarkPopCount/PopCount_use_iter_one_bit_by_one_bit
// BenchmarkPopCount/PopCount_use_iter_one_bit_by_one_bit
// BenchmarkPopCount/PopCount_use_iter_one_bit_by_one_bit-16
// 1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
// === RUN   BenchmarkPopCount/PopCount_use_n_&=_n-1
// BenchmarkPopCount/PopCount_use_n_&=_n-1
// BenchmarkPopCount/PopCount_use_n_&=_n-1-16
// 1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
// PASS
// ok      github.com/Suuu775/gopl/ch2/ex2_3       0.368s
