package communication

import (
	"elevator"
	"elevdriver"
	"fmt"
)

func ReadOrderButtons() {
	var (
		floor int
		dir   elevdriver.Direction
		eDir  elevator.Direction
	)
	fmt.Println("Comm.input--> Started reading orders...")
	for {
		floor, dir = elevdriver.GetButton()
		fmt.Println("Comm.input--> Recived order")

		// Need to cast type because of the redeclaration of type
		// 'Direction' and 'Order' in package 'elevator'
		eDir = elevator.Direction(dir)
		switch eDir {
		case elevator.NONE:
			commChan.intOrder <- elevator.Order{floor, eDir}
			//fmt.Println("Comm.input--> Internal order sent")
		case elevator.UP, elevator.DOWN:
			commChan.extOrder <- elevator.Order{floor, eDir}
			//fmt.Println("Comm.input--> External order sent")
		default:
			fmt.Println("Comm.input--> Discarding corrupt order")
		}
	}
}
