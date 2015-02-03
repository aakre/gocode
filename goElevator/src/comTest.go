package main

import(
	"fmt"
	"communication"
	"elevator"
)

func main(){

	order := elevator.Order{3, elevator.NONE}
	communication.CodeTST(order)
	fmt.Println("comTest finished")
	//communication.StartSystem()

}
