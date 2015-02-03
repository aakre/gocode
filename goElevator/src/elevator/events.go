package elevator

import (
	"elevdriver"
	"fmt"
	"time"
)

type Event int

const (
	STOP Event = iota
	WAIT
	GO_UP
	GO_DOWN
	EMERG
	OBSTRUCTION
)

func (elev *elevator) getEvent() (event Event) {
	switch {
	case <-intChan.emergency:
		return EMERG
	case <-intChan.obstruction:
		return OBSTRUCTION
	}
	event = WAIT
	atFloor := elevdriver.AtFloor()
	switch elev.state {
	case BOOT:
		return
	case IDLE, STOPPED:
		// Decide to stay/go based on elev.orders.
		if atFloor {
			event = elev.eventDirection()
		} else {
			event = EMERG
			fmt.Println("Elevator--> Slipped past floor")
			fmt.Println("Elevator--> Consider adjusting 'BACKLASHCONST' in 'motorcontrol'")
		}
	case ASCENDING, DECENDING:
		// If at floor, decide to stop or not based on elev.orders.
		// Since the elevator is moving, we check the floor again inside eventStop().
		event = elev.eventStop()
	case EMERGENCY:
		// Decide to exit emergency state or not
		event = elev.eventEmergency()
	}
	return
}

func (elev *elevator) handleEvent(event Event) {
	switch elev.state {
	case BOOT:
		elev.statemachineBoot()
	case IDLE:
		elev.statemachineIdle(event)
	case STOPPED:
		elev.statemachineStopped(event)
	case ASCENDING:
		elev.statemachineAscending(event)
	case DECENDING:
		elev.statemachineDecending(event)
	case EMERGENCY:
		elev.statemachineEmergency(event)
	}
}

func (elev *elevator) eventStop() (event Event) {
	event = WAIT

	currFloor := elevdriver.GetFloor()
	if currFloor != -1 {
		elev.lastFloor = currFloor
	} else {
		return
	}

	switch {
	case currFloor == 1 && elev.state == DECENDING:
		// Stop if at bottom floor and decending
		event = STOP
	case currFloor == NFLOORS && elev.state == ASCENDING:
		// Stop if at top floor and ascending
		event = STOP
	case len(elev.orders) == 0:
		// Stop if no orders left
		event = STOP
	case elev.findOrder(elev.suitableOrder()) != -1:
		// Stop if suitable orders at this floor
		event = STOP
	case !elev.ordersLeftInThisDir():
		event = STOP
	}
	return
}

func (elev *elevator) eventDirection() (event Event) {
	event = WAIT

	if len(elev.orders) == 0 {
		return
	}

	switch {
	case elev.orders[0].Floor == elev.lastFloor:
		// If next order is at current floor, stop (open doors)
		event = STOP
		elev.lastDir = NONE
	case elev.findOrder(elev.suitableOrder()) != -1:
		// Stop (open door) if other suitable orders at current floor
		event = STOP
	case elev.orders[0].Floor > elev.lastFloor:
		// If next order is above, go up
		event = GO_UP
	case elev.orders[0].Floor < elev.lastFloor:
		// If next order is below, go down
		event = GO_DOWN
	}
	return
}

func (elev *elevator) eventEmergency() (event Event) {
	// Exits 'EMERGENCY' when (simulated) maintenance is given
	event = WAIT
	select {
	case <-intChan.maintenance:
		if elevdriver.AtFloor() {
			event = STOP
		} else {
			event = GO_DOWN
		}
	default:
	}
	return
}

func eventMaintenance() {
	// Simulating maintenance: OBS on, STOP, STOP, OBS off
	fmt.Println("Elevator--> Waiting for maintenance...")
	for {
		if <-intChan.obstruction && <-intChan.emergency {
			fmt.Println("Elevator--> Maintenance in progress...")
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	for {
		if !<-intChan.obstruction && !<-intChan.emergency {
			fmt.Println("Elevator--> Maintenance complete")
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	intChan.maintenance <- true
}
