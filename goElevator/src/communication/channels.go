package communication

import (
	"elevator"
	"network"
	"systemstate"
)

type CommChannels struct {
	intOrder      chan elevator.Order
	extOrder      chan elevator.Order
	extCommand    chan elevator.Order
	extInform     chan systemstate.StateOrder
	extComplete   chan systemstate.StateOrder
	newOffer      chan Offer
	stateUpdate   chan systemstate.SysState
	updateRequest chan string
	resync        chan bool
	panic         chan bool
}

var commChan CommChannels
var elevChan elevator.ElevChannels
var stateChan systemstate.StateChannels
var netChan network.NetChannels

func commChanInit() {
	commChan.intOrder = make(chan elevator.Order)
	commChan.extOrder = make(chan elevator.Order)
	commChan.extCommand = make(chan elevator.Order)
	commChan.extInform = make(chan systemstate.StateOrder)
	commChan.extComplete = make(chan systemstate.StateOrder)
	commChan.newOffer = make(chan Offer)
	commChan.stateUpdate = make(chan systemstate.SysState)
	commChan.updateRequest = make(chan string)
	commChan.resync = make(chan bool)
	commChan.panic = make(chan bool)
}

func initChannels() {
	commChanInit()
	elevChan.ElevChanInit()
	stateChan.StateChanInit()
	netChan.NetChanInit()
}
