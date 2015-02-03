package main

import (
		"fmt"
		"crypto/tls"
		"os"
	
)

const(
	encUser = "b3l2aW5kLmFha3Jl"
	encPass = "bWFubmUwMTA1"
	defaultUser = "oyvind.aakre@gmail.com"
)

type mail struct{
	sendTo string
	from string
	subject string
	body string
}

var responseCh chan string

func main() {
	//email:= composeMail()
	email:= mail{defaultUser, defaultUser, "blank", "some message"}
	fmt.Println("Okay! Now ready to send your email!")
	tls_sock, err := tls.Dial("tcp", "smtp.gmail.com:465", nil)
	fmt.Println("Dial error: ", err)
	fmt.Println("Dial up successfull...")
	responseCh = make(chan string)
	go serverResponse(tls_sock)
	fmt.Println("Authenticating...")
	auth(tls_sock)
	fmt.Println("Sending...")
	sendEmail(email, tls_sock)
	fmt.Println("Email successfully sent!")
}

func serverResponse(tls_sock *tls.Conn) {
	send := true
	for {
		var buf [512]byte
		numBytes, err := tls_sock.Read(buf[0:])
		fmt.Println(numBytes, " bytes read. Error: ", err)
		response := string(buf[0:numBytes])
		responseCode := string(buf[0:3])
		if err != nil {
			fmt.Println("Response from server, error: ", err)
		}
		fmt.Println("Response: ", response)
		if send {
			responseCh <- responseCode
			send = false
		}
	}
}

func auth(tls_sock *tls.Conn){
	response := <- responseCh
	validate(response, "220")
	
	n, err := tls_sock.Write([]byte("EHLO"))
	fmt.Println(n, err)
	//response = <- responseCh
	
	n, err = tls_sock.Write([]byte("auth login"))
	fmt.Println(n, err)
	//response = <- responseCh
	
	n, err = tls_sock.Write([]byte(encUser))
	fmt.Println(n, err)
	//response = <- responseCh
	
	
	n, err = tls_sock.Write([]byte(encPass))
	fmt.Println(n, err)
	//response = <- responseCh
}

func sendEmail(email mail, tls_sock *tls.Conn) {
	n, err2 := tls_sock.Write([]byte("mail from: <" + defaultUser + ">"))
	fmt.Println(n, err2)
	
	n, err2 = tls_sock.Write([]byte("rcpt to: <" + email.sendTo + ">"))
	fmt.Println(n, err2)
	
	n, err2 = tls_sock.Write([]byte("data"))
	fmt.Println(n, err2)
	
	n, err2 = tls_sock.Write([]byte("subject: <" + email.subject + ">"))
	fmt.Println(n, err2)
	
	n, err2 = tls_sock.Write([]byte(""))
	fmt.Println(n, err2)
	n, err2 = tls_sock.Write([]byte(email.body))
	fmt.Println(n, err2)
	n, err2 = tls_sock.Write([]byte(""))
	fmt.Println(n, err2)
	n, err2 = tls_sock.Write([]byte("."))
	fmt.Println(n, err2)
	
	
	n, err2 = tls_sock.Write([]byte("quit"))
	fmt.Println(n, err2)
	tls_sock.Close()
}

func composeMail() mail {
	fmt.Println("Send mail to: ")
	sendTo := singleReadInput()
	fmt.Println("Subject: ")
	subject := singleReadInput()
	fmt.Println("Ready to read message. End message with '$' \nMessage: ")
	var message, input string
	done := false
	for !done {
		input = singleReadInput()
		if input == "$" {
			done = true
		} else {
			message += input + " "
		}
	}
	email := mail{sendTo, defaultUser, subject, message}
	return email	
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

func validate(response, code string) {
	if response != code {
		fmt.Println("Error validating response code ", code)
		os.Exit(1)
	} else {
		fmt.Println("Valid code ", code)
	}
}