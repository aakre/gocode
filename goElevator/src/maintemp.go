package main

import (
	"communication"
	"fmt"
)

func main() {
	fmt.Println("Main--> Starting system...")
	communication.OperateSystem()
	fmt.Println("Main--> System shutdown")
}
