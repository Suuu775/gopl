package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep, indexs := "", "", ""
	for i, arg := range os.Args[1:] {
		indexs += sep + fmt.Sprint(i)
		s += sep + arg
		sep = " "
	}
	fmt.Println(indexs)
	fmt.Println(s)
}
