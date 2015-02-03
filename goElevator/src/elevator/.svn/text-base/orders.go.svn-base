package elevator

import (
	"fmt"
)

// Same type as in 'elevdriver', redeclared to shorten code in this package
type Order struct {
	Floor int
	Dir   Direction
}

func (elev *elevator) addOrder(newOrder Order) {
	// Append new order to order queue
	if !elev.duplOrder(newOrder) {
		elev.orders = append(elev.orders, newOrder)
		fmt.Println("Elevator--> New order added:", newOrder)
	} else {
		fmt.Println("Elevator--> Duplicate order", newOrder, "not added")
	}
}

func (elev *elevator) findOrder(searchOrder Order) (orderIndex int) {
	// Check if we have suitable orders in the order queue
	orderIndex = -1

	for i, order := range elev.orders {
		if searchOrder.Floor == order.Floor {
			if order.Dir == NONE || searchOrder.Dir == NONE || searchOrder.Dir == order.Dir {
				orderIndex = i
				return
			}
		}
	}
	return
}

func (elev *elevator) duplOrder(searchOrder Order) (duplicate bool) {
	// Check for duplicate order in queue
	duplicate = false
	for _, order := range elev.orders {
		if searchOrder.Floor == order.Floor && searchOrder.Dir == order.Dir {
			duplicate = true
		}
	}
	return
}

func (elev *elevator) serveOrders(serveOrder Order) (del bool) {
	// Complete and delete orders suitable for serveOrder
	del = false
	for {
		i := elev.findOrder(serveOrder)
		if i != -1 {
			go completeOrder(elev.orders[i])
			elev.deleteOrder(i)
			del = true
		} else {
			return
		}
	}
	return
}

func (elev *elevator) delExtOrders() {
	for i := 0; i < len(elev.orders); {
		switch elev.orders[i].Dir {
		case UP, DOWN:
			elev.deleteOrder(i)
		default:
			i++
		}
	}
	fmt.Println("Elevator--> Deleted external orders")
}

func (elev *elevator) deleteOrder(index int) {
	elev.orders = append(elev.orders[:index], elev.orders[index+1:]...)
}

func completeOrder(order Order) {
	// Inform com module that an order has been completed
	fmt.Println("Elevator---> Completed order:", order)
	extChan.CplOrder <- order
	//fmt.Println("Elevator---> cplOrder sent")
}

func (elev *elevator) ordersLeftInThisDir() (ordersLeft bool) {
	ordersLeft = false
	switch elev.lastDir {
	case UP:
		for _, order := range elev.orders {
			if order.Floor > elev.lastFloor {
				ordersLeft = true
				return
			}
		}
	case DOWN:
		for _, order := range elev.orders {
			if order.Floor < elev.lastFloor {
				ordersLeft = true
				return
			}
		}
	}
	return
}

func (elev *elevator) printOrders() {
	var dir string
	if len(elev.orders) > 0 {
		fmt.Println("***ORDERS***")
		for i, val := range elev.orders {
			switch val.Dir {
			case UP:
				dir = "UP"
			case DOWN:
				dir = "DOWN"
			case NONE:
				dir = "NONE"
			}
			fmt.Printf("[ #%d | F: %d | D: %s ]\n", i, val.Floor, dir)
		}
	} else {
		fmt.Println("***NO ORDERS***")
	}
}

func (elev *elevator) suitableOrder() (order Order) {
	// Returns an order that suits the elevators current state.
	// Used to compare with available orders etc.
	return Order{elev.lastFloor, elev.lastDir}
}
