package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type Communication struct {
	Name    string
	Message string
}

func handleRequest(conn net.Conn) {
	var msg string
	fmt.Println("handles start")
	dec := gob.NewDecoder(conn)
	for msg != "End" {
		err := dec.Decode(&msg)

		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("Recieved", msg)
		}
	}

	conn.Close()
}

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:3333")
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	go handleRequest(conn)
	name, _ := reader.ReadString('\n')

	for {

		msg, _ := reader.ReadString('\n')
		comm := Communication{name, msg}
		//Scanf seems to be reading a second blank line
		//for every input.

		// Creating a new encoder doesn't seem the best
		// way to do this.
		err := gob.NewEncoder(conn).Encode(&comm)

		if err != nil {
			fmt.Println("2", err)
		}

	}
}
