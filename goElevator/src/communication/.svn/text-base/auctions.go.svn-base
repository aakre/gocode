package communication

import (
	"elevator"
	"fmt"
	"systemstate"
	"time"
)

type Offer struct {
	IP    string
	Cost  int
	Order elevator.Order
}

func (sys *system) runAuction(order elevator.Order) {
	fmt.Println("comm.Auctions--> Starting auction of order:", order)

	elevChan.GetCost <- order
	bestCost := <-elevChan.SendCost
	fmt.Println("comm.Auctions--> My cost:", bestCost)
	bestElevator := "me"

	timeout := false
	var passedTime time.Duration
	startTime := time.Now()
	go sendNEW(order)
	fmt.Println("comm.Auctions--> Waiting for offers...")
	for !timeout {
		select {
		case newOffer := <-commChan.newOffer:
			if newOffer.Order == order {
				fmt.Println("comm.Auctions--> Received new offer from", newOffer.IP, ":", newOffer.Cost)
				if newOffer.Cost < bestCost {
					bestCost = newOffer.Cost
					bestElevator = newOffer.IP
				}
			} else {
				fmt.Println("comm.Auctions--> Old offer received from", newOffer.IP, ". Ignoring...")
			}
			passedTime = time.Now().Sub(startTime)
			//fmt.Println("comm.Auctions--> Auction passed time:", passedTime)
		case <-time.After(AUCTIONDURATION*time.Millisecond - passedTime):
			timeout = true

			/*if time.Since(startTime) > (AUCTIONTIMEOUT * time.Millisecond){
				timeout = true
			}*/
		}
	}
	if bestElevator == "me" {
		if sys.operational {
			fmt.Println("comm.Auctions--> I won the auction for", order)
			elevChan.NewOrder <- order
		} else {
			fmt.Println("comm.Auctions--> Elevator not operational, but no other offers received.")
			fmt.Println("comm.Auctions--> Saving order,", order, ", for later execution...")
		}
		stateChan.AddOrder <- systemstate.StateOrder{sys.myIP, order}
		go sendINF(order)
	} else {
		fmt.Println("comm.Auctions-->", bestElevator, "won the auction for", order)
		go sendCMD(bestElevator, order)
	}
}

func (sys *system) auctionOrders(elevIP string) {
	fmt.Println("comm.Auctions--> Starting auctions for", elevIP, "'s external orders...")
	stateChan.FindExtOrders <- elevIP
	orders := <-stateChan.OrdersFound
	if len(orders) > 0 {
		fmt.Println("comm.Auctions--> Auctioning orders:", orders)
		for _, order := range orders {
			sys.runAuction(order)
		}
		fmt.Println("comm.Auctions--> Auctions for", elevIP, "complete")
	} else {
		fmt.Println("comm.Auctions-->", elevIP, "had no orders")
	}
}
