//------I'm Alive Module (IMA)------//

package network

import (
	"fmt"
	"net"
	"time"
)

func imaListen() {
	service := bcast + ":" + UDPport
	addr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		fmt.Println("network.IMAListen()--> ResolveUDP error")
		internalChan.setupfail <- true
	}
	sock, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println("network.IMAListen()--> ListenUDP error")
		internalChan.setupfail <- true
	}
	var data [512]byte
	for {
		select {
		case <-internalChan.quitImaListen:
			return
		default:
			_, remoteAddr, err := sock.ReadFromUDP(data[0:])
			if err != nil {
				fmt.Println("network.IMAListen()--> ReadFromUDP error")
				break
			}
			if LOCALIP != remoteAddr.IP.String() {
				if err == nil {
					elevIP := remoteAddr.IP.String()
					internalChan.ima <- elevIP
				} else {
					fmt.Println("network.IMAListen()--> UDP read error")
				} //End if
			} //End if
		} //End select
	} //End for
} //End imaListen

func imaWatcher() {
	peers := make(map[string]time.Time)
	deadline := IMALOSS * IMAPERIOD * time.Millisecond
	for {
		select {
		case ip := <-internalChan.ima:
			_, inMap := peers[ip]
			if inMap {
				peers[ip] = time.Now()
			} else {
				peers[ip] = time.Now()
				internalChan.newIP <- ip
			}
		case <-time.After(ALIVEWATCH * time.Millisecond):
			for ip, timestamp := range peers {
				if time.Now().After(timestamp.Add(deadline)) { // elevator missed deadline
					fmt.Println("network.IMAWatcher--> Timeout ", ip)
					externalChan.GetDeadElevator <- ip //notify AI of dead elevator
					internalChan.closeConn <- ip       //notify TCPMap
					delete(peers, ip)
				}
			}
		case deadIP := <-externalChan.SendDeadElevator: //FROM decode DTH
			internalChan.closeConn <- deadIP
			delete(peers, deadIP)
		case errorIP := <-internalChan.errorIP:
			_, inMap := peers[errorIP]
			if inMap {
				externalChan.Panic <- true
			}
		case <-internalChan.quitImaWatcher:
			return
		}
	}
}

func imaSend() {
	service := bcast + ":" + UDPport
	addr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		fmt.Println("network.IMASend()--> Resolve error")
		internalChan.setupfail <- true
	}
	imaSock, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Println("network.IMASend()--> Dial error")
		internalChan.setupfail <- true
	}
	ima := []byte("IMA")
	for {
		select {
		case <-internalChan.quitImaSend:
			return
		default:
			_, err := imaSock.Write(ima)
			if err != nil {
				fmt.Println("network.IMASend()--> UDP send error")
			}
			time.Sleep(IMAPERIOD * time.Millisecond)
		}
	}
}
