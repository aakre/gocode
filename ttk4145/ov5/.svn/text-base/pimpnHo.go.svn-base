package main

import(
	"fmt"
	"net"
	"os"
	"time"
	"encoding/json"
)

const IMAPER = 100 	//Period between each IMA in milliseconds
const IMALOSS = 5	//Number of missing IMA to be interpeded as DEAD
const PRINTINT = 1000	//Number of milliseconds between printing of value

func main(){
	//Initialize and define sockets etc
	service := "localhost:9001"
	
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	
	udpConn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	
	lastVal := Ho(udpConn)
	udpConn.Close()
		
	udpConn, err = net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	
	Pimp(udpConn, lastVal)
}

func Ho(udpConn net.Conn) (lastVal int){
	lastVal = 0
	pimpExist := false
	valChan := make(chan int)
	
	go func(udpConn net.Conn, valChan chan int){
		var buf [512]byte
		var v int
		for{
			n, _ := udpConn.Read(buf[0:])
			//checkError(err)
			
			_ = json.Unmarshal(buf[0:n], &v)
			valChan<- v
		}
	}(udpConn, valChan)
	
	fmt.Println("Ho--> Listening for Pimps...")
	
	for{
		select{
			case lastVal = <-valChan:
				pimpExist = true
			case <-time.After(IMALOSS * IMAPER * time.Millisecond):
				if pimpExist{
					fmt.Println("Ho--> Pimp is dead, Imma takeover dis shit...")
				}else{
					fmt.Println("Ho--> No pimps around, Imma take dis shit...")
				}
				return lastVal
		}
	}
	return lastVal
}

func Pimp(udpConn net.Conn, lastVal int){
	valChan := make(chan int)
	
	go func(udpConn net.Conn, valChan chan int){
		var v int
		var err error
		for{
			select{
				case v = <-valChan:
					buf, err := json.Marshal(&v)
					checkError(err)
					_, _ = udpConn.Write(buf)
				default:
					buf, err := json.Marshal(&v)
					checkError(err)
					_, _ = udpConn.Write(buf)
			}
			time.Sleep(IMAPER * time.Millisecond)
			checkError(err)
		}
	}(udpConn, valChan)
	
	fmt.Println("***What's happening?***")
	fmt.Println("***~Dun, dun, dundundun~***")
	fmt.Println("***Your Ho has evolved into a Pimp!***")
	
	for{
		lastVal ++
		fmt.Println("Pimp--> Value is: ", lastVal)
		valChan<- lastVal
		time.Sleep(PRINTINT * time.Millisecond)		
	}
	
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
