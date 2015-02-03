package network

import (
	"fmt"
	"math/rand"
	"time"
)

func StartNetwork(NetChan NetChannels) {
	rand.Seed(time.Now().UTC().UnixNano()) // Seed used when attempting to connect
	internalChan.init()
	externalChan = NetChan
	go imaWatcher()
	go imaListen()
	go imaSend()
	go manageTCPConnections()
	for {
		select {
		case <-internalChan.setupfail:
			fmt.Println("***net.Startup--> Setupfail. Retrying...")
			internalChan.quitImaSend <- true
			internalChan.quitImaListen <- true
			internalChan.quitImaWatcher <- true
			internalChan.quitListenTCP <- true
			internalChan.quitTCPMap <- true
			time.Sleep(time.Millisecond) //Sleep to let functions finish
			go imaListen()
			go imaSend()
			go manageTCPConnections()
		case <-time.After(NETSETUP * time.Millisecond):
			fmt.Println("net.Startup--> Network setup complete")
			return
		}
	}
}
