package main
//chmod +x 'name'

//tune the log handling
import (
		"fmt"
		"net"
		"time"
		"encoding/json"
		"bufio"
		"os"
		"strings"
		"io/ioutil"
)

type sessionState int

const(
	LOGIN sessionState = iota
	IDLE
	CHAT
	PREPARE_PRIVATE
	ACCEPT_MSG
	QUIT
	DISCONNECTED
)

type Message struct {
	Sender string
	Receiver string
	Payload string
}



type chatSession struct{
	server string
	nick string
	conn net.Conn
	state sessionState
	userInput chan string
	serverMsg chan Message
	currentChat string
	serverStatus chan string
	log map[string]string
	lostConnection chan bool
}

const(
	port = ":31337"
	defaultService = "78.91.27.136" + port
	WRITEDL = 100
	CONNATMPT = 5
	DIALINT = 100
	VERSION = "v1.3"
	CHATPATH = "chat_history/"
	DIALTIMEOUT = 200*time.Millisecond
)

func connect(service string) net.Conn {
	attempts := 0
	for attempts < CONNATMPT {		
		conn, err := net.DialTimeout("tcp4", service, DIALTIMEOUT)
		if err != nil {
			//fmt.Println("connect---> DialTCP error when connecting to", service)
			attempts++
			time.Sleep(DIALINT * time.Millisecond)
		} else {
			//fmt.Println("connect---> Connection made to server")
			return conn
			break
		}
	}
	return nil
}


func main () {
	session := chatSession{
		conn: nil,
		state: LOGIN,
		currentChat: "NONE",
		userInput: make(chan string),
		serverMsg: make(chan Message),
		serverStatus: make(chan string),
		lostConnection: make(chan bool),
		log: make(map[string]string),
	}
	clearScreen()
	startScreen()
	ready := false
	for !ready {
		session.printMenu()
		server := session.requestServer()
		session.conn = connect(server)
		if session.conn == nil {
			fmt.Println("Sorry, no connection could be made to", server)
			fmt.Println("Do you want to try again (y/n)?")
			reply := singleReadInput()
			if reply == "n" {
				fmt.Println("OK. Quitting application...")
				os.Exit(0)
			} else {
				clearScreen()
			}
		} else {
			ready = true
		}
	}
	//err := os.MkdirAll(CHATPATH, 2777)
	//fmt.Println(err)
	clearScreen()
	welcomeScreen()
	session.registerUser()
	session.state = IDLE
	clearScreen()
	session.printMenu()
	go session.readFromServer()
	go session.contReadInput()
	for {
		select {
		case input:= <- session.userInput:
			session.handleInput(input)
		case msg := <- session.serverMsg:
			session.handleServerMsg(msg)
		case <-session.lostConnection:
			session.saveChat()
			clearScreen()
			session.state = DISCONNECTED
			session.printMenu()
		}
	}	
}


func (s *chatSession) printMenu() {
	switch s.state {
	case IDLE:
		fmt.Println("Currently in chat with", s.currentChat)
		fmt.Println("Available commands:")
		fmt.Println("'q public' for a public session \n'q private' for a private session \n'q status' to see other users on server \n'q menu' to bring up this menu")
	case CHAT:
		fmt.Println("Currently in chat with", s.currentChat)
		fmt.Println("Available commands:")
		fmt.Println("Simply write to send a message \n'q leave' to leave session \n'q status' to see other users on server \n'q menu' to bring up this menu")
	case LOGIN:
		fmt.Println("\nPlease select a server:")
		fmt.Printf("Type 'd' for default server: %s\n",	defaultService)
		fmt.Println("...or enter an IP adress you think might be a Nostalgica server")	
	case DISCONNECTED:
		fmt.Println("Seems like we lost connection to the server")
		fmt.Println("Shall I try to reconnect to ", s.server, "? (y/n)")
	}
}

func (s *chatSession) requestServer() string {
	server := singleReadInput()
	if server == "d" {
		server = defaultService
	} else {
		server = server + port
	}
	s.server = server
	return server
}

func (s *chatSession) registerUser() {
	complete := false
	for !complete {
		fmt.Println("Please enter your preferred username:")
		nick := singleReadInput()
		fmt.Println("Checking if nick is available...")
		ok := s.checkAvailable(nick)
		if ok {
			complete = true
			s.nick = nick
			fmt.Println("You have been registered at the server with username", nick)
		} else {
			fmt.Println("Ooops! Seems like ", nick, "is already taken...")
		}
	}
}

func (s *chatSession)checkAvailable(nick string) bool {
	b := []byte(nick)
	header := []byte("REG")
	b = append(header,b...)
	s.conn.Write(b)
	var data [512]byte
	numbytes, err := s.conn.Read(data[0:])
	if err == nil && string(data[0:3]) == "REG" {
		if string(data[3:numbytes]) == "OK" {
			return true
		}
	}
	return false
}

func (s *chatSession) contReadInput() {
	rd := bufio.NewReader(os.Stdin)
	for {
		line, err := rd.ReadString('\n')
		if err == nil {
			line = strings.TrimRight(line, "\n")
			s.userInput <- line
		}
	}
}

func singleReadInput() string {
	var input string
	for n, err := fmt.Scanf("%s", &input); err== nil; n,err = fmt.Scanf("%s", &input) {
		if n!=0 && err== nil {
			return input
		}
	}
	fmt.Println("Something went wrong when reading your precious keystrokes... :'(")
	return ""
}

func (s *chatSession) readFromServer() {
	for {
		var data [512]byte
		numbytes, err := s.conn.Read(data[0:])
		if err == nil {
			header := string(data[0:3])
			switch header{
			case "MSG": //Public and private Messages goes here
				var msg Message
				err := json.Unmarshal(data[3:numbytes], &msg)
				if err != nil {
					fmt.Println("PANIC PANIC PANIC UNMARSHAL ERROR!", err)
				} else {
					s.serverMsg <- msg
				}
			case "STS":
				var m Message
				err := json.Unmarshal(data[3:numbytes], &m)
				if err == nil {
					go func() {
						s.serverStatus <- m.Payload
					}()
				} else {
					fmt.Println("PANIC PANIC PANIC UNMARSHAL ERROR!", err)
				}
			}
		} else {
			fmt.Println("Connection error", err)
			s.lostConnection <- true
			return
		}
	}
}

func (s *chatSession) handleInput(input string) {
	switch s.state {
	case CHAT:
		s.handleInputChat(input)
	case IDLE:
		s.handleInputIdle(input)
	case PREPARE_PRIVATE:
		s.preparePrivate(input)
	case ACCEPT_MSG:
		s.acceptMsg(input)
	case DISCONNECTED:
		s.handleInputDisconnected(input)
	}
}

func (s *chatSession) handleInputIdle(input string) {
	switch input{
		case "q public", "q pub", "q publ", " publi": //User wants to join public chat
			clearScreen()
			s.togglePublic()
			s.currentChat = "PUBLIC"
			s.loadChatHistory()
			s.displayChatHistory()
			s.state = CHAT
			s.printMenu()
		case "q private", "q pri", "q priv", "q privat":
			clearScreen()
			//Provide list of users to send private Message
			fmt.Println("To whom do you want to send a message?")
			s.getServerStatus(true, "")
			s.state = PREPARE_PRIVATE
		case "q status", "q stat", "q sta":
			s.getServerStatus(true, "")
		case "q menu", "q meny", "q men":
			clearScreen()
			s.printMenu()
		case "q quit", "q exit":
			clearScreen()
			fmt.Println("NOSTALGICA: Quitting application...")
			s.state = QUIT
			s.saveChat()
			s.disconnect()
			os.Exit(0)
		case "q":
			s.printMenu()
		default:
			clearScreen()
			fmt.Println("Sorry! Input not recognized: ", input)
			s.printMenu()
	}
}

func (s *chatSession) preparePrivate(input string) {
	if input == s.nick {
		fmt.Println("Shizofrenic alert! Trying to chat privately with yourself")
	} else {
		validName, _ := s.getServerStatus(false, input)
		if validName {
			s.currentChat = input
			s.state = CHAT
			clearScreen()
			s.printMenu()
			s.loadChatHistory()
			s.displayChatHistory()
		} else {
			clearScreen()
			fmt.Println("NOSTALGICA: Username not recognized...", input)
			s.state = IDLE
			s.printMenu()
		}
	}
}

func (s *chatSession) handleInputChat(input string) {
	switch input {
	case "q leave", "q l", "q back":
		fmt.Println("NOSTALGICA: Leaving chat...")
		if s.currentChat == "PUBLIC" { s.togglePublic() }
		s.saveChat()
		s.currentChat = "NONE"
		s.state = IDLE
		clearScreen()
		s.printMenu()
	case "q exit", "q e", "q quit":
		clearScreen()
		fmt.Println("NOSTALGICA: Quitting application...")
		s.state = QUIT
		s.saveChat()
		s.disconnect()
		os.Exit(0)
	case "q status", "q stat", "q statu", "q s":
		s.getServerStatus(true, "")
	case "q menu", "q men", "q meny", "q m":
		s.printMenu()
	case "q private", "q priv":
			clearScreen()
			s.saveChat()
			//Provide list of users to send private Message
			fmt.Println("To whom do you want to send a message?")
			printStatus := true
			s.getServerStatus(printStatus, "")
			s.state = PREPARE_PRIVATE
			if s.currentChat == "PUBLIC" { 
				s.togglePublic() 
			}
	case "q public", "q pub":
		if s.currentChat != "PUBLIC" {
			s.saveChat()
			fmt.Println("NOSTALGICA: Leaving chat...")
			s.togglePublic()
			s.currentChat = "PUBLIC"
			s.state = CHAT
			clearScreen()
			s.printMenu()
			s.loadChatHistory()
			s.displayChatHistory()
		} else {
			fmt.Println("Silly you;) You are already in the public chat!")
		}
	case "q":
		s.printMenu()
	default: 
		m := Message{Sender: s.nick, Receiver: s.currentChat, Payload: input}
		b, err := json.Marshal(m)
		if err == nil {
			header := []byte("MSG")
			b = append(header,b...)
			s.conn.Write(b)
		} else {
			fmt.Println("NOSTALGICA: Something went wrong. Do you mind trying again?")
			s.printMenu()
		}
	}
}

func (s *chatSession) handleInputDisconnected (input string) {
	switch input {
	case "y": //try to reconnect
		s.conn = connect(s.server)
		if s.conn == nil {
			clearScreen()
			s.printMenu()
		} else {
			ok := s.checkAvailable(s.nick)
			if ok {
				fmt.Println("Connection to server reestablished...")
				time.Sleep(3*time.Second)
				s.state = IDLE
				clearScreen()
				s.printMenu()
			} else {
				clearScreen()
				s.printMenu()
			}
		}
	case "n": //quit
		fmt.Println("Quitting application...")
		os.Exit(0)
	default:
		fmt.Println("NOSTALGICA: Input not recognized...", input)
	}

}

func (s *chatSession) getServerStatus(print bool, checkName string) (bool, string) {
	header := []byte("STS")
	s.conn.Write(header)
	status := <-s.serverStatus
	if print { printStatus(status) }
	if checkName != "\n" {
		inList := searchList(status,checkName)
		return inList, status
	}
	return true, status
}


func (s *chatSession) togglePublic() {
	header := []byte("PUB")
	msg := Message{Payload: s.nick}
	b, _ := json.Marshal(msg)
	b = append(header, b...)
	s.conn.Write(b)
}

func (s *chatSession) disconnect() {
	header := []byte("BYE")
	msg := Message{Payload: s.nick}
	b, _ := json.Marshal(msg)
	b = append(header, b...)
	s.conn.Write(b)
}


func (s *chatSession) handleServerMsg(msg Message) {
	switch s.state {
	case IDLE:
		s.handleMsgIdle(msg)
	case CHAT:
		s.handleMsgChat(msg)
	case ACCEPT_MSG:
		s.updateLog(msg)
		fmt.Println("----> log updated in handleserverMsg")
	}
}

func (s *chatSession) acceptMsg(input string) {
	switch input {
	case "y", "yes":
		s.saveChat()
		s.state = CHAT
		s.loadChatHistory()
		s.displayChatHistory()
	case "n", "no":
		s.state = IDLE
		s.currentChat = "NONE"
		clearScreen()
		s.printMenu()
	default:
		fmt.Println("Okay then. Be like that...")
	}
}

func (s *chatSession) handleMsgIdle(msg Message) {
	clearScreen()
	s.state = ACCEPT_MSG
	s.currentChat = msg.Sender
	fmt.Printf("\nNOSTALGICA: Message received from %s!\n", msg.Sender)
	fmt.Printf("%q\n", msg.Payload)
	s.updateLog(msg)
	fmt.Println("----> log updated in handleMsgIdle")
	fmt.Println("Do you want to go to chat now? (y/n)")
}

func (s *chatSession) handleMsgChat(msg Message) {
	if msg.Sender == s.currentChat || msg.Receiver == s.currentChat{ //Handle both private and public msg in one check
		if msg.Sender != s.nick {
			fmt.Printf("%s: %s\n",msg.Sender, msg.Payload)
		}
	} else {
		fmt.Printf("\nNOSTALGICA: Message received from %s!\n", msg.Sender)
	}
	s.updateLog(msg)
	fmt.Println("----> log updated in handlMsgChat")
}

func startScreen() {
	fmt.Printf("\nWelcome to NOSTALGICA 2013 %s!\n", VERSION)
	fmt.Println("\nNOSTALGICA is a GUI-free chat client running in the terminal!")
	fmt.Println("Cut the cord, throw away the mouse and get ready to use the keyboard!")
}

func welcomeScreen() {
	fmt.Println("Lucky bastard! You have successfully connected to a server.")
	fmt.Println("No it is time to register!")
}



func clearScreen() {
	for i:=0; i<100; i++ {
		fmt.Println("")
	}
}

func printStatus(users string) {
	fmt.Println("\nNOSTALGICA: Status of connected users...")
	start := 0
	for i:=0; i<len(users); i++ {
		if users[i] == '!' {
			fmt.Println(users[start:i]," -- status: is active in public chat")
			start = i+1
		} else if users[i] == '?' {
			fmt.Println(users[start:i]," -- status: I dunno... Try sending a message maybe?")
			start = i+1
		}	
	}
	fmt.Println("")
}

func searchList(users, checkName string) bool {
	start := 0
	for i:=0; i<len(users); i++ {
		if users[i] == '!' || users[i] == '?'{
			if users[start:i] == checkName { return true }
			start = i+1
		}
	}
	return false
}

func (s *chatSession) updateLog(msg Message) {
	oldLog, noLog := s.log[msg.Sender]
	if noLog {
		/*fmt.Println("--->updateLog: No live log found, trying to load from file...")
		if s.loadChatHistory() {
			fmt.Println("--->updateLog: Was able to load previous log from file and update it")
			oldLog = s.log[msg.Sender]
			s.log[msg.Sender] = oldLog + fmt.Sprintf("%s: %s\n",msg.Sender, msg.Payload)
		} else {*/
			fmt.Println("--->updateLog: No live log found")
			s.log[msg.Sender] = fmt.Sprintf("%s: %s\n",msg.Sender, msg.Payload)
			fmt.Println("--->updateLog: Log created...")
		//}
	} else {
		fmt.Println("--->updateLog: Live log found and updated")
		s.log[msg.Sender] = oldLog + fmt.Sprintf("%s: %s\n",msg.Sender, msg.Payload)
	}
}

func (s *chatSession) saveChat() {
	file := s.currentChat + ".txt"
	b := []byte(s.log[s.currentChat])
	err := ioutil.WriteFile(file, b, 0644)
    if err != nil { fmt.Println(err) }
    //delete(s.log, s.currentChat)
}

func (s *chatSession) loadChatHistory() bool {
	file := s.currentChat + ".txt"
	b, err := ioutil.ReadFile(file)
    if err == nil && len(b)>0{
    	//fmt.Printf("\nWe saved your latest chat history with %s\n", s.currentChat)
    	//fmt.Println(string(b))
    	s.log[s.currentChat] = string(b)
    	fmt.Println("--->chat history loaded from file...")
    	return true
    }
    fmt.Println("---> could not load chat history from file...")
	fmt.Println(err)
	return false
}

func (s *chatSession) displayChatHistory() {
	fmt.Println(s.log[s.currentChat])
}

func GetMyIP() string {
	allIPs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("GetMyIP(): Error receiving IPs. IP set to localhost. Consider setting IP manually")
		return "localhost"
	}

	IPString := make([]string, len(allIPs))
	for i := range allIPs {
		temp := allIPs[i].String()
		ip := strings.Split(temp, "/")
		IPString[i] = ip[0]
	}
	var myIP string
	for i := range IPString {
		if IPString[i][0:3] == "129" {
			myIP = IPString[i]
		}
	}
	return myIP
}






