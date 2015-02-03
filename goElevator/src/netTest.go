package main

import (
	"fmt"
	"network"
)

func main() {
	var (
		comCh       = make(chan string)
		connect     = make(chan string)
		connectFail = make(chan string) //channel for feedback from connectTcp-function
	)
	//Sets up elevatorMap and a function that controls access to this map
	updateImaMapCh := network.ProtectImaMapInit(comCh, connect, connectFail)
	go network.IMAListen(comCh, updateImaMapCh) //Listen for other elevators
	ip := <-comCh
	fmt.Println("death: elevator with ip ", ip, " has died")
}
