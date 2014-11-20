package main

import ("fmt"
"net"
"encoding/gob")

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

type Person struct{
	Name string
	Age int
}

func main(){
	l, err := net.Listen(CONN_TYPE, 
		CONN_HOST + ":" + CONN_PORT)
		if err != nil{
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
        fmt.Println("Handles Incoming")
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  var msg Person
  
  for{
  err := gob.NewDecoder(conn).Decode(&msg)
  
  if err!= nil{
	fmt.Println(err)
	return
  }else{
	fmt.Println("Recieved", msg)}
	response := "Hello " + msg.Name
	gob.NewEncoder(conn).Encode(response)
  }
  
  conn.Close()
}


