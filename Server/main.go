package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

type Person struct {
	Name string
	Conn net.Conn
}

type Communication struct {
	Name    string
	Message string
}

var people [2]Person = [2]Person{}
var index = 0

func main() {
	l, err := net.Listen(CONN_TYPE,
		CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	defer l.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}

		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	var msg Communication
	// First message should be the username
	gob.NewDecoder(conn).Decode(&msg)
	people[index] = Person{msg.Name, conn}
	index += 1

	for {
		err := gob.NewDecoder(conn).Decode(&msg)

		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("Recieved", msg)
		}

		for _, v := range people {
			if v.Name != msg.Name {
				gob.NewEncoder(v.Conn).Encode(msg)
			}

		}
	}

	conn.Close()
}
