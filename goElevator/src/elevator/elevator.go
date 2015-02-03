package elevator

import (
	"elevdriver"
	"fmt"
	"time"
)

// Same type as in 'elevdriver', redeclared to shorten code in this package
type Direction int

const (
	UP Direction = iota
	DOWN
	NONE
)

type elevator struct {
	state      State
	orders     []Order
	lastFloor  int
	lastDir    Direction
	doorOpen   bool
	doorTimer  time.Time
	obsTimeout bool
	obsTimer   time.Time
}

func HandleElevator(elevChan ElevChannels) {
	// Initialization
	intChan.init()
	extChan = elevChan
	var elev elevator
	elev.elevInit()

	// Running for-loop
	for {
		select {
		case newOrder := <-extChan.NewOrder:
			elev.addOrder(newOrder)

		case costOrder := <-extChan.GetCost:
			cost := elev.calculateCost(costOrder)
			extChan.SendCost <- cost

		default:
		}
		//elev.printOrders()
		//elev.printStatus(event)
		event := elev.getEvent()
		elev.handleEvent(event)
		time.Sleep(ELEVLOOPINT * time.Millisecond)
	}
}

func (elev *elevator) elevInit() {
	fmt.Println("Elevator--> Initializng...")
	elevdriver.Init()

	elev = &elevator{
		state:     BOOT,
		lastFloor: -1,
		lastDir:   NONE,
		doorOpen:  false,
		doorTimer: time.Now(),
	}

	go emergWatcher()
	go obsWatcher()
	go floorWatcher()
	fmt.Println("Elevator--> Initialization complete")
}

func (elev *elevator) printStatus(event Event) {
	var s1, s2, s3 string
	switch elev.state {
	case BOOT:
		s1 = "BOOT"
	case IDLE:
		s1 = "IDLE"
	case STOPPED:
		s1 = "STOPPED"
	case ASCENDING:
		s1 = "ASCENDING"
	case DECENDING:
		s1 = "DECENDING"
	case EMERGENCY:
		s1 = "EMERGENCY"
	}
	switch event {
	case STOP:
		s2 = "STOP"
	case WAIT:
		s2 = "WAIT"
	case GO_UP:
		s2 = "GO UP"
	case GO_DOWN:
		s2 = "GO DOWN"
	case EMERG:
		s2 = "EMERG"
	case OBSTRUCTION:
		s2 = "OBSTRUCTION"
	}
	switch elev.lastDir {
	case NONE:
		s3 = "NONE"
	case UP:
		s3 = "UP"
	case DOWN:
		s3 = "DOWN"
	}
	fmt.Printf("Elevator--> State: %s Event: %s LastFloor: %d LastDir: %s\n",
		s1, s2, elev.lastFloor, s3)
}
