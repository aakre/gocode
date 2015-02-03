package elevator

import (
	"elevdriver"
	"time"
)

func (elev *elevator) motorUp() {
	elevdriver.MotorUp()
}

func (elev *elevator) motorDown() {
	elevdriver.MotorDown()
}

func (elev *elevator) motorStop() {
	// Implements active breaking (reversing motor) when stopping
	elevdriver.MotorStop()
	switch elev.state {
	case ASCENDING:
		elevdriver.MotorDown()
	case DECENDING, BOOT:
		elevdriver.MotorUp()
	}
	time.Sleep(BACKLASHCONST * time.Millisecond)
	elevdriver.MotorStop()
}
