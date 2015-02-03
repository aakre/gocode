package elevator

import (
	"elevdriver"
	"fmt"
	"time"
)

func (elev *elevator) closeDoor() bool {
	// Close door, if able, and handle if obstruction timed out 
	if time.Now().After(elev.doorTimer.Add(DOORTIMER * time.Second)) {
		elevdriver.ClearDoor()
		elev.doorOpen = false
		elev.handleDoorObstruction(false)
		fmt.Println("Elevator--> Closing door...")
		return true
	}
	return false
}

func (elev *elevator) doorReset() {
	elev.doorTimer = time.Now()
}

func (elev *elevator) openDoor() {
	elev.doorTimer = time.Now()
	elevdriver.SetDoor()
	elev.doorOpen = true
	fmt.Println("Elevator--> Opening door...")
}

func (elev *elevator) handleDoorObstruction(obs bool) {
	switch obs {
	case true:
		elev.doorReset()
		if (elev.obsTimer == time.Time{}) {
			// Obstruction is new; set timer
			elev.obsTimer = time.Now()
		} else {
			// Obstruction for too long triggers measures to divert orders (like emergency)
			if !elev.obsTimeout && time.Now().After(elev.obsTimer.Add(OBSTIMER*time.Second)) {
				fmt.Println("Elevator--> Obstruction timeout")
				elev.obsTimeout = true
				extChan.Operational <- false
				elev.delExtOrders()
			}
		}
	case false:
		// Reset timer, and if obstruction timed out; reset and inform AI
		elev.obsTimer = time.Time{}
		if elev.obsTimeout {
			elev.obsTimeout = false
			extChan.Operational <- true
		}
	}
}
