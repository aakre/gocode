package main

import (
	//"fmt"
	"network"
	"time"
	"communication"
)


func main() {
	var netChan network.NetChannels
	netChan.NetChanInit()
	network.NetSetup(netChan)
	time.Sleep(1*time.Second)
	go repeater(netChan)
	//time.Sleep(9*time.Second)
	for {
		
		newMail:= <- netChan.Inbox
		communication.Decode(newMail)
		
	}
	/*
	fmt.Println(string(newMail.Msg))
	deadIP :=<- netChan.GetDeadElevator
	fmt.Println("netChan: deadIP ", deadIP)
	*/
}

func repeater(netChan network.NetChannels){
	for {
		mail := network.Mail{Msg: []byte("TST")}
		netChan.SendToAll <- mail
		time.Sleep(2*time.Second)
	}
}
