package main

import (
	"fmt"
	"math/rand"
	"time"
)

func protect(x *int, write, read chan int) {
	for {
		select {
		case v := <-write:
			*x += v
			fmt.Println("Var x changed to: ", *x, "\n")
		case read <- *x:
		}
	}
}

func readVar(read chan int) {
	for i:= range read{
		fmt.Println("-->The value of x is: ",i, "\n")
		time.Sleep(time.Second)
	}
}

func add(ch chan int) {
	i := 0
	for {
		ch <- 2*i
		i++
		randSleep := 1000 + rand.Intn(3000)
		time.Sleep(time.Duration(randSleep) * time.Millisecond)
	}
}

func subtr(ch chan int) {
	i := 5
	for {
		ch <- -i
		i++
		randSleep := 1000 + rand.Intn(2000)
		time.Sleep(time.Duration(randSleep) * time.Millisecond)
	}
}


func main() {
	varChan := make(chan int)
	readChan := make(chan int)
	x := 0
	go protect(&x, varChan, readChan)
	go add(varChan)
	go subtr(varChan)
	go readVar(readChan)

	for {
		time.Sleep(1 * time.Second)
	}
}
