package main

import (
		//"os"
		//"bufio"
		//"io"
		"io/ioutil"
		"fmt"
)

func main() {
	// read the whole file
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
    	fmt.Println("No such file")
    }

    // write file body
    err = ioutil.WriteFile("output.txt", b, 0644)
    if err != nil { panic(err) }
}