package systemstate

import (
	"elevator"
	"fmt"
)

type SysState struct {
	Internal map[string][]bool
	External [][]string
}

type StateOrder struct {
	Id    string
	Order elevator.Order
}

func HandleSysState(stateChan StateChannels) {
	fmt.Println("SysState--> Initializing...")
	extChan = stateChan
	state := SysState{}
	state.init()
	fmt.Println("SysState--> Initialization complete")

	for {
		//state.print()
		select {
		case newState := <-extChan.SysState:
			state.overwrite(newState)
			fmt.Println("SysState--> Overwritten state")

		case <-extChan.SendState:
			extChan.SysState <- state.copy()
			//fmt.Println("SysState--> State sent")

		case sOrder := <-extChan.CheckOrder:
			extChan.OrderExists <- state.checkOrder(sOrder)
			//fmt.Println("SysState--> Order checked")

		case sOrder := <-extChan.AddOrder:
			state.addOrder(sOrder)
			//fmt.Println("SysState--> Added order")

		case sOrder := <-extChan.DelOrder:
			state.delOrder(sOrder)
			//fmt.Println("SysState--> Removed order:", sOrder)

		case elevId := <-extChan.FindIntOrders:
			extChan.OrdersFound <- state.intOrdersWithId(elevId)
			//fmt.Println("SysState--> Internal orders found")

		case elevId := <-extChan.FindExtOrders:
			extChan.OrdersFound <- state.extOrdersWithId(elevId)
			//fmt.Println("SysState--> External orders found")
		}
	}
}

func (state *SysState) init() {
	state.Internal = make(map[string][]bool)
	state.External = make([][]string, elevator.NFLOORS)
	for i := range state.External {
		state.External[i] = make([]string, 2)
	}
}

func (state *SysState) print() {
	fmt.Println("----System State----")
	fmt.Println("**Internal Orders**")
	for id, s := range state.Internal {
		fmt.Println("ID:", id, "| Orders:", s)
	}
	fmt.Println("**External Orders**")
	for floor, dir := range state.External {
		fmt.Println("Floor:", floor, "| UP:", dir[0], "| DOWN:", dir[1])
	}
	fmt.Println("--------------------")
}
