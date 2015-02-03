package main

import(
	"fmt"
	"net"
	"os"
	"time"
	"encoding/json"
)

func main(){
	service := ":9001"

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err, true)

	listenSoc, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, true)
	fmt.Println("Listening...")

	newConnChan := make(chan net.Conn)
	closeConnChan := make(chan net.Conn)
	go clientLibrary(closeConnChan, newConnChan)

	for{
		connSoc, err := listenSoc.Accept()
		if err == nil{
			go handleClient(connSoc, closeConnChan)
			fmt.Println("2")
			newConnChan <- connSoc
			fmt.Println("3")
		}else{
			checkError(err, false)
		}

		time.Sleep(100 * time.Millisecond)
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

func handleClient(clientSoc net.Conn, closeConnChan chan net.Conn){
	var buf [512]byte
	var numBytes int
	var err error

	clientAddr := clientSoc.RemoteAddr()
	localAddr := clientSoc.LocalAddr()
	fmt.Printf("New client from %s connected at %s \n", clientAddr.String(), localAddr.String())

	for{
		fmt.Println("0")
		numBytes, err = clientSoc.Read(buf[0:])
		fmt.Println("1")
		if err != nil{
			fmt.Println("Error while reading from client\nClosing connection...")
			clientSoc.Close()
			return
		}

		fmt.Printf("Message with length %d recived from client %s\n", numBytes, clientAddr.String())
		fmt.Printf("---> %s \n", string(buf[0:numBytes]))
	}
}

func clientLibrary(closeConnChan, newConnChan chan net.Conn){
	var clientSlice []net.Conn

	for{
		newConn := <-newConnChan
		clientSlice = append(clientSlice,newConn)
		fmt.Println("Currently connected clients:")
		for i := range clientSlice{
			fmt.Printf("[ #%d | %s ]\n", i, clientSlice[i].RemoteAddr())


			buf, err := json.Marshal(&clientSlice)
			checkError(err, false)
			if err == nil{
				_, err = clientSlice[i].Write(buf)
				checkError(err, false)
			}
		}
	}
}







