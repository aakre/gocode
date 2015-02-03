package elevator

import (
	"elevdriver"
	"time"
)

func emergWatcher() {
	emergency := false
	for {
		select {
		case intChan.emergency <- emergency:
		default:
			emergency = elevdriver.GetStopButton()
		}
		time.Sleep(WATCHINTa * time.Millisecond)

	}
}

func obsWatcher() {
	obstruction := false
	for {
		select {
		case intChan.obstruction <- obstruction:
		default:
			obstruction = elevdriver.GetObs()
		}
		time.Sleep(WATCHINTa * time.Millisecond)

	}
}

func floorWatcher() {
	floor := -1
	for {
		floor = elevdriver.GetFloor()
		if floor != -1 {
			elevdriver.SetFloor(floor)
		}
		time.Sleep(WATCHINTb * time.Millisecond)
	}
}
