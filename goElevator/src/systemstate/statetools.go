package systemstate

import (
	"elevator"
)

func (state *SysState) copy() (stateCopy SysState) {
	stateCopy.init()
	for key, sli := range state.Internal {
		stateCopy.Internal[key] = make([]bool, elevator.NFLOORS)
		for i, v := range sli {
			stateCopy.Internal[key][i] = v
		}
	}
	for flr, dir := range state.External {
		stateCopy.External[flr][0] = dir[0]
		stateCopy.External[flr][1] = dir[1]
	}
	return
}

func (state *SysState) overwrite(newState SysState) {
	state.init()
	for key, sli := range newState.Internal {
		state.Internal[key] = make([]bool, elevator.NFLOORS)
		for i, v := range sli {
			state.Internal[key][i] = v
		}
	}
	for flr, dir := range newState.External {
		state.External[flr][0] = dir[0]
		state.External[flr][1] = dir[1]
	}
}

func (state *SysState) checkOrder(sOrder StateOrder) (exists bool) {
	exists = false
	switch sOrder.Order.Dir {
	case elevator.NONE:
		intOrd, inMap := state.Internal[sOrder.Id]
		if inMap {
			exists = intOrd[sOrder.Order.Floor-1]
		}
	case elevator.UP, elevator.DOWN:
		exists = state.External[sOrder.Order.Floor-1][sOrder.Order.Dir] != ""
	}
	return
}

func (state *SysState) addOrder(sOrder StateOrder) {
	switch sOrder.Order.Dir {
	case elevator.NONE:
		_, inMap := state.Internal[sOrder.Id]
		if !inMap {
			state.Internal[sOrder.Id] = make([]bool, elevator.NFLOORS)
		} else {
			// Error handling since we have been asked to remove a non-existant order?
			// This would imply that our sysState isn't up to date etc.
			// However, during testing, it could simply be caused by hard-coded orders.
		}
		state.Internal[sOrder.Id][sOrder.Order.Floor-1] = true
	case elevator.UP, elevator.DOWN:
		state.External[sOrder.Order.Floor-1][sOrder.Order.Dir] = sOrder.Id
	}
}

func (state *SysState) delOrder(sOrder StateOrder) {
	switch sOrder.Order.Dir {
	case elevator.NONE:
		_, inMap := state.Internal[sOrder.Id]
		if inMap {
			state.Internal[sOrder.Id][sOrder.Order.Floor-1] = false
		} else {
			extChan.Resync <- true
		}
	case elevator.UP, elevator.DOWN:
		if state.External[sOrder.Order.Floor-1][sOrder.Order.Dir] != "" {
			state.External[sOrder.Order.Floor-1][sOrder.Order.Dir] = ""
		} else {
			extChan.Resync <- true
		}
	}
}

func (state *SysState) intOrdersWithId(elevId string) (orders []elevator.Order) {
	intOrders, inMap := state.Internal[elevId]
	if inMap {
		for floor, orderAtThisFloor := range intOrders {
			if orderAtThisFloor {
				orders = append(orders, elevator.Order{floor + 1, elevator.NONE})
			}
		}
	}
	return
}

func (state *SysState) extOrdersWithId(elevId string) (orders []elevator.Order) {
	for floor, dir := range state.External {
		if dir[0] == elevId {
			orders = append(orders, elevator.Order{floor + 1, elevator.UP})
		}
		if dir[1] == elevId {
			orders = append(orders, elevator.Order{floor + 1, elevator.DOWN})
		}

	}
	return
}
