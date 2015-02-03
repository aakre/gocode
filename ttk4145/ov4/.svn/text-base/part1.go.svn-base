package main

import (
	"fmt"
	"time"
)

type Square struct {
	side int
}

func (s Square) Area() (area int) {
	area = s.side * s.side
	return
}

func print() {
	for {
		fmt.Println("Hello! I am a gorutine. You like penis! HAHA!\n")
		time.Sleep(time.Second)
	}
}

func f1(c chan Square) {
	x:=0
	for {
		c <- Square{x}
		x++
		time.Sleep(time.Second)
	}
}

func f2(c chan Square) {
	var y Square
	for {
		y = <-c
		fmt.Println(y.Area(),"\n")
	}
}


func main() {
	sq1 := Square{5}
	sq2 := Square{9}
	array := []Square{sq1, sq2}

	for _, value := range array {
		fmt.Println("Area of square is ", value.Area(), "\n")
	}

	
	
	i := 0
	for i < 5{
		go print() // Kan lage sånn ca. ganske mange
		fmt.Println(i)
		i++
	}
	
	
	ch := make(chan Square)
	go f1(ch)
	go f2(ch)
	
	never := make(chan int)
	<-never
}



