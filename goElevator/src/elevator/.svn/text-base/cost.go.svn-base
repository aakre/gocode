package elevator

import (
	//"fmt"
	"elevdriver"
)

func (elev *elevator) calculateCost(order Order) (cost int) {
	dirCost, waitCost, travCost, stateCost := 0, 0, 0, 0
	queueLength := len(elev.orders)
	switch elev.state {
	case BOOT:
		stateCost = FLOORWEIGHT * (NFLOORS - 1)
	case IDLE:
		cost = getTravelCost(elev.lastFloor, order.Floor)
		return
	case STOPPED:
		if elev.obsTimeout {
			cost = OBSCOST
			return
		}
		if getTravelCost(elev.lastFloor, order.Floor) == 0 {
			cost = 0
			return
		}
		switch elev.lastDir {
		case UP:
			dirCost, waitCost, travCost = elev.costAscending(order)
		case DOWN:
			dirCost, waitCost, travCost = elev.costDecending(order)
		case NONE:
		}
	case ASCENDING:
		if order.Dir == UP && order.Floor == elevdriver.GetFloor() {
			cost = QUEUEWEIGHT * queueLength
			return
		}
		dirCost, waitCost, travCost = elev.costAscending(order)
	case DECENDING:
		if order.Dir == DOWN && order.Floor == elevdriver.GetFloor() {
			cost = QUEUEWEIGHT * queueLength
			return
		}
		dirCost, waitCost, travCost = elev.costDecending(order)
	case EMERGENCY:
		cost = 9001
		return
	}

	cost = dirCost + waitCost + travCost + stateCost
	return
}

func (elev *elevator) costAscending(order Order) (dirCost, waitCost, travCost int) {
	if orderIsAbove(elev.lastFloor, order) {
		travCost = getTravelCost(elev.lastFloor, order.Floor)
		waitCost = elev.ordersBetween(elev.lastFloor, UP, order) * STOPWEIGHT
	} else {
		dirCost = DIRWEIGHT
		lastFloor := elev.currOrdersLastFloor()
		travCost = getTravelCost(elev.lastFloor, lastFloor)
		travCost += getTravelCost(lastFloor, order.Floor)
		waitCost = elev.ordersBetween(elev.lastFloor, UP, order) * STOPWEIGHT
		waitCost += elev.ordersBetween(lastFloor, DOWN, order) * STOPWEIGHT
	}
	return
}

func (elev *elevator) costDecending(order Order) (dirCost, waitCost, travCost int) {
	if orderIsBelow(elev.lastFloor, order) {
		travCost = getTravelCost(elev.lastFloor, order.Floor)
		waitCost = elev.ordersBetween(elev.lastFloor, DOWN, order) * STOPWEIGHT
	} else {
		dirCost = DIRWEIGHT
		lastFloor := elev.currOrdersLastFloor()
		travCost = getTravelCost(elev.lastFloor, lastFloor)
		travCost += getTravelCost(lastFloor, order.Floor)
		waitCost = elev.ordersBetween(elev.lastFloor, DOWN, order) * STOPWEIGHT
		waitCost += elev.ordersBetween(lastFloor, UP, order) * STOPWEIGHT
	}
	return
}

func getTravelCost(floor int, orderFloor int) (travCost int) {
	diff := floor - orderFloor
	switch {
	case diff == 0:
		travCost = 0
	case diff < 0:
		travCost = (-diff) * FLOORWEIGHT
	case diff > 0:
		travCost = diff * FLOORWEIGHT
	}
	return
}

func orderIsAbove(floor int, costOrder Order) bool {
	switch costOrder.Dir {
	case UP, NONE:
		return costOrder.Floor > floor
	}
	return false
}

func orderIsBelow(floor int, costOrder Order) bool {
	switch costOrder.Dir {
	case DOWN, NONE:
		return costOrder.Floor < floor
	}
	return false
}

func (elev *elevator) ordersBetween(floor int, dir Direction, costOrder Order) (numOrders int) {
	numOrders = 0
	if floor == costOrder.Floor {
		return
	}
	switch dir {
	case UP:
		for _, order := range elev.orders {
			for flr := floor; flr < costOrder.Floor; flr++ {
				if order.Floor == flr && (order.Dir == dir || order.Dir == NONE) {
					numOrders++
				}
			}
		}
	case DOWN:
		for _, order := range elev.orders {
			for flr := floor; flr > costOrder.Floor; flr-- {
				if order.Floor == flr && (order.Dir == dir || order.Dir == NONE) {
					numOrders++
				}
			}
		}
	}
	return
}

func (elev *elevator) currOrdersLastFloor() (floor int) {
	floor = elev.lastFloor
	if len(elev.orders) == 0 {
		return
	}
	switch elev.lastDir {
	case UP:
		for _, order := range elev.orders {
			if order.Floor > floor {
				floor = order.Floor
			}
		}
	case DOWN:
		for _, order := range elev.orders {
			if order.Floor < floor {
				floor = order.Floor
			}
		}
	}
	return
}
