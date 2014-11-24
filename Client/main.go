package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type Person struct {
	Name string
	Age  int
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

	for {
		msg, _ := reader.ReadString('\n')

		//Scanf seems to be reading a second blank line
		//for every input.

		person := Person{msg, 20}

		// Creating a new encoder doesn't seem the best
		// way to do this.
		err := gob.NewEncoder(conn).Encode(&person)

		if err != nil {
			fmt.Println("2", err)
		}

	}
}
