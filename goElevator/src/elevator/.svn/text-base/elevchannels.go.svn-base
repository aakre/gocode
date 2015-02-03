package elevator

type internalChannels struct {
	emergency   chan bool
	obstruction chan bool
	maintenance chan bool
}

type ElevChannels struct {
	NewOrder    chan Order
	CplOrder    chan Order
	GetCost     chan Order
	SendCost    chan int
	Operational chan bool
}

var intChan internalChannels
var extChan ElevChannels

func (intChan *internalChannels) init() {
	intChan.emergency = make(chan bool)
	intChan.obstruction = make(chan bool)
	intChan.maintenance = make(chan bool)
}

func (elevChan *ElevChannels) ElevChanInit() {
	elevChan.NewOrder = make(chan Order)
	elevChan.CplOrder = make(chan Order)
	elevChan.GetCost = make(chan Order)
	elevChan.SendCost = make(chan int)
	elevChan.Operational = make(chan bool)
}
