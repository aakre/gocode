package elevator

import (
	"elevdriver"
	"fmt"
)

type State int

const (
	BOOT State = iota
	IDLE
	STOPPED
	ASCENDING
	DECENDING
	EMERGENCY
)

func (elev *elevator) statemachineBoot() {
	// Elevator decends until it gets to a floor and stops
	// Then it sets the correct states and stops
	fmt.Println("Elevator--> Booting...")
	elev.motorDown()
	elev.lastDir = DOWN
	for !elevdriver.AtFloor() {
	}
	elev.motorStop()
	elev.lastFloor = elevdriver.GetFloor()
	elev.state = IDLE
	elev.lastDir = NONE
	fmt.Println("Elevator--> Boot complete")
}

func (elev *elevator) statemachineIdle(event Event) {
	switch event {
	case STOP:
		elev.serveOrders(elev.suitableOrder())
		elev.openDoor()
		elev.state = STOPPED
	case WAIT:
	case GO_UP:
		elev.motorUp()
		elev.state = ASCENDING
		elev.lastDir = UP
	case GO_DOWN:
		elev.motorDown()
		elev.state = DECENDING
		elev.lastDir = DOWN
	case EMERG:
		elev.handleEmergency()
	case OBSTRUCTION:
		elev.openDoor()
		elev.state = STOPPED
	}
}

func (elev *elevator) statemachineStopped(event Event) {
	switch event {
	case STOP:
		elev.serveOrders(elev.suitableOrder())
		elev.doorReset()
	case WAIT:
		if elev.closeDoor() {
			elev.state = IDLE
			elev.lastDir = NONE
		}
	case GO_UP:
		if elev.closeDoor() {
			elev.motorUp()
			elev.state = ASCENDING
			elev.lastDir = UP
		}
	case GO_DOWN:
		if elev.closeDoor() {
			elev.motorDown()
			elev.state = DECENDING
			elev.lastDir = DOWN
		}
	case EMERG:
		elev.handleEmergency()
	case OBSTRUCTION:
		elev.handleDoorObstruction(true)
	}
}

func (elev *elevator) statemachineAscending(event Event) {
	switch event {
	case STOP:
		elev.motorStop()
		elev.serveOrders(elev.suitableOrder())
		elev.openDoor()
		elev.state = STOPPED
	case WAIT:
	case GO_UP:
	case GO_DOWN:
	case EMERG:
		elev.motorStop()
		elev.handleEmergency()
	case OBSTRUCTION:
		elev.motorStop()
		elev.handleEmergency()
	}
}

func (elev *elevator) statemachineDecending(event Event) {
	switch event {
	case STOP:
		elev.motorStop()
		elev.serveOrders(elev.suitableOrder())
		elev.openDoor()
		elev.state = STOPPED
	case WAIT:
	case GO_UP:
	case GO_DOWN:
	case EMERG:
		elev.motorStop()
		elev.handleEmergency()
	case OBSTRUCTION:
		elev.motorStop()
		elev.handleEmergency()
	}
}

func (elev *elevator) statemachineEmergency(event Event) {
	switch event {
	case STOP:
		extChan.Operational <- true
		elevdriver.ClearStopButton()
		elev.doorReset()
		elev.serveOrders(elev.suitableOrder())
		elev.state = STOPPED
	case WAIT:
		if elev.doorOpen && !elevdriver.AtFloor() {
			elev.closeDoor()
		}
	case GO_UP:
		extChan.Operational <- true
		elevdriver.ClearStopButton()
		if !elev.doorOpen || elev.closeDoor() {
			elev.motorUp()
			elev.state = ASCENDING
			elev.lastDir = UP
		} else {
			elev.state = STOPPED
		}
	case GO_DOWN:
		extChan.Operational <- true
		elevdriver.ClearStopButton()
		if !elev.doorOpen || elev.closeDoor() {
			elev.motorDown()
			elev.state = DECENDING
			elev.lastDir = DOWN
		} else {
			elev.state = STOPPED
		}
	case EMERG:
	case OBSTRUCTION:
	}
}

func (elev *elevator) handleEmergency() {
	// Inform AI (diverts orders and ignores new ones) and wait for maintenance
	fmt.Println("Elevator--> Emergency")
	extChan.Operational <- false
	go eventMaintenance()
	if elevdriver.AtFloor() {
		elev.openDoor()
	} else {
		elev.closeDoor()
	}
	elev.state = EMERGENCY
	elevdriver.SetStopButton()
	elev.delExtOrders()
}
