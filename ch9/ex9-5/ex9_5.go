package ex95

import (
	"fmt"
	"time"
)

func Ex95(){
	in := make(chan int)
	out := make(chan int)

	go pingpong(in,out)
	go pingpong(out,in)
	in<-0
	time.Sleep(1*time.Second)
	fmt.Println(<-in)
}

func pingpong(in chan int,out chan int){
	for{
		out <- (<-in+1)
	}
}
