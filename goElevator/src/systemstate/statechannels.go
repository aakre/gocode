package systemstate

import (
	"elevator"
)

type StateChannels struct {
	CheckOrder    chan StateOrder
	OrderExists   chan bool
	AddOrder      chan StateOrder
	DelOrder      chan StateOrder
	SendState     chan bool
	SysState      chan SysState
	FindIntOrders chan string
	FindExtOrders chan string
	OrdersFound   chan []elevator.Order
	Resync        chan bool
}

var extChan StateChannels

func (stateChan *StateChannels) StateChanInit() {
	stateChan.CheckOrder = make(chan StateOrder)
	stateChan.OrderExists = make(chan bool)
	stateChan.AddOrder = make(chan StateOrder)
	stateChan.DelOrder = make(chan StateOrder)
	stateChan.SendState = make(chan bool)
	stateChan.SysState = make(chan SysState)
	stateChan.FindIntOrders = make(chan string)
	stateChan.FindExtOrders = make(chan string)
	stateChan.OrdersFound = make(chan []elevator.Order)
	stateChan.Resync = make(chan bool)
}
