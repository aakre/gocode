package main

import(
	"fmt"
	"net"
	"os"
	"time"
	"encoding/json"
)

func main(){
	host := "localhost:9001"
	
	tcpAddr, err := net.ResolveTCPAddr("tcp", host)
	checkError(err, true)

	connSoc, err := net.DialTCP("tcp", nil, tcpAddr)
	for err != nil{
		fmt.Println("Could not connect to server, retrying...")
		time.Sleep(2 * time.Second)
		connSoc, err = net.DialTCP("tcp", nil, tcpAddr)
	}
	fmt.Println("Connected to server")

	go userList(connSoc)
	fmt.Println("0")
	
	for{
		fmt.Println("1")
		sendBuf := []byte("I am alive!")
		fmt.Println("2")
		
		_, err = connSoc.Write(sendBuf[0:])
		checkError(err, true)
		fmt.Println("IAA message sent")

		time.Sleep(3 * time.Second)
		
	}
}

func checkError(err error, exit bool) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		if exit{
			fmt.Println("Exiting...")
			os.Exit(1)
		}
	}
}

func userList(connSoc net.Conn){
	var rcvBuf[1024]byte
	var numBytes int
	var err error

	var clientSlice []net.Conn

	fmt.Println("a")
	
	for{
		numBytes, err = connSoc.Read(rcvBuf[0:])
		fmt.Println("\nb\n")
		checkError(err, false)
		err = json.Unmarshal(rcvBuf[0:numBytes], &clientSlice)
		fmt.Println("\nc\n")
		checkError(err, false)
		
		fmt.Println("d")
		if err == nil {
			fmt.Println("Currently connected users:")
			for i := range clientSlice{
				fmt.Printf("[#%d | %s ]\n", i, clientSlice[i].RemoteAddr())
			}
		}else{
			fmt.Println(err.Error())
		}
		time.Sleep(100 * time.Millisecond)
	}
}
