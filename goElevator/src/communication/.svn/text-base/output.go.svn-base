package communication

import (
	"elevdriver"
	"fmt"
	"network"
	"time"
)

func SetOrderLights() {
	// Retrives a copy of system state and sets lights accordingly
	myIp := network.GetMyIP()
	fmt.Println("Comm.output--> Started setting of order lights...")
	for {
		//fmt.Println("Comm.output--> Asking for state...")
		stateChan.SendState <- true
		//fmt.Println("Comm.output--> Waiting for state...")
		currState := <-stateChan.SysState
		//fmt.Println("Comm.output--> State recived")
		intOrd, inMap := currState.Internal[myIp]
		if inMap {
			for flr, val := range intOrd {
				if val {
					elevdriver.SetLight(flr+1, elevdriver.NONE)
				} else {
					elevdriver.ClearLight(flr+1, elevdriver.NONE)
				}
			}
		}

		for flr, dir := range currState.External {
			if dir[elevdriver.UP] != "" {
				elevdriver.SetLight(flr+1, elevdriver.UP)
			} else {
				elevdriver.ClearLight(flr+1, elevdriver.UP)
			}
			if dir[elevdriver.DOWN] != "" {
				elevdriver.SetLight(flr+1, elevdriver.DOWN)
			} else {
				elevdriver.ClearLight(flr+1, elevdriver.DOWN)
			}
		}
		time.Sleep(LIGHTSETINT * time.Millisecond)
	}
}
