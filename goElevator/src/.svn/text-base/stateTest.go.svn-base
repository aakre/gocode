package main

import (
	"elevator"
	"fmt"
	"systemstate"
	"time"
)

func main() {

	stateChan := systemstate.StartSysState()

	newOrder := systemstate.StateOrder{
		Id:    "10.10.10.10",
		Order: elevator.Order{2, elevator.NONE},
	}
	newOrder2 := systemstate.StateOrder{
		Id:    "10.10.10.10",
		Order: elevator.Order{1, elevator.UP},
	}
	newOrder3 := systemstate.StateOrder{
		Id:    "10.10.10.10",
		Order: elevator.Order{1, elevator.DOWN},
	}

	stateChan.AddOrder <- newOrder
	var tempState systemstate.SysState
	stateChan.SendState<- true
	tempState = <-stateChan.SysState
	fmt.Println("tempState:", tempState)
	
	stateChan.CheckOrder<- newOrder
	fmt.Println(<-stateChan.OrderExists)
	stateChan.CheckOrder<- newOrder2
	fmt.Println(<-stateChan.OrderExists)
	
	stateChan.AddOrder <- newOrder2
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("tempState:", tempState)
	stateChan.SysState<- tempState
	
	stateChan.AddOrder <- newOrder3
	time.Sleep(100 * time.Millisecond)
	fmt.Println("tempState:", tempState)

	never := make(chan int)
	<-never
}
