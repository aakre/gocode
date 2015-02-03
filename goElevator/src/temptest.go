package main

import (
	"elevator"
	"elevdriver"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Startup...")
	//elevdriver.Init()
	elevChan := elevator.StartElevator()
	go complOrders(elevChan)
	fmt.Println("Penis")

	//printButton wont work if we dont wait for elevdriver.Init() to run
	time.Sleep(time.Second)
	elevChan.GetCost<- elevator.Order{3, elevator.NONE}
	fmt.Println("Cost:", <-elevChan.SendCost)
	elevChan.NewOrder<- elevator.Order{3, elevator.NONE}

	//go printFloor()
	go printButton()
	//go printObs()
	//go printEmerg()
	time.Sleep(10 * time.Second)
	elevChan.NewOrder<- elevator.Order{1, elevator.NONE}
	
	
	never := make(chan int)
	<-never
}

func printFloor() {
	for {
		fmt.Println("YOLO")
		fmt.Println("Floor: ", elevdriver.GetFloor())
	}
}

func printButton() {
	fmt.Println("Buttonprinter started...")
	for {
		floor, dir := elevdriver.GetButton()
		fmt.Println("Button (F,D): ", floor, dir)
	}
}

func printObs() {
	for {
		fmt.Println("Obs: ", elevdriver.GetObs())
	}
}

func printEmerg() {
	for {
		elevdriver.GetStopButton()
		fmt.Println("Emergency")
	}
}

func complOrders(elevChan elevator.ElevChannels){
	for{
		completedOrder := <-elevChan.CplOrder
		fmt.Println("Elevator has completed order:", completedOrder)
	}
}
