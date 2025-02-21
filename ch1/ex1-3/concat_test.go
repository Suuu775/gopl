// Running tool: D:\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkEcho$ github.com/Suuu775/gopl/ch1/ex1-3

// goos: windows
// goarch: amd64
// pkg: github.com/Suuu775/gopl/ch1/ex1-3
// cpu: 12th Gen Intel(R) Core(TM) i5-12500H
// === RUN   BenchmarkEcho
// BenchmarkEcho
// === RUN   BenchmarkEcho/Inefficient_echo_Benchmark
// BenchmarkEcho/Inefficient_echo_Benchmark
// BenchmarkEcho/Inefficient_echo_Benchmark-16
// 1000000000               0.0000097 ns/op               0 B/op          0 allocs/op
// === RUN   BenchmarkEcho/Efficient_echo_Benchmark
// BenchmarkEcho/Efficient_echo_Benchmark
// BenchmarkEcho/Efficient_echo_Benchmark-16
// 1000000000               0.0000023 ns/op               0 B/op          0 allocs/op
// PASS
// ok      github.com/Suuu775/gopl/ch1/ex1-3       1.473s

package ex13_test

import (
	"os"
	"strings"
	"testing"
)

var args = []string{"0", "1", "2", "hello", "world", "blg", "t1"}

func BenchmarkEcho(b *testing.B) {
	b.Run("Inefficient_echo Benchmark", func(b *testing.B) {
		inefficient_echo()
	})

	b.Run("Efficient_echo Benchmark", func(b *testing.B) {
		efficient_echo()
	})
}

func inefficient_echo() {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
}

func efficient_echo() {
	strings.Join(args[1:], " ")
}
