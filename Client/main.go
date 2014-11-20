package main

import ( "fmt"
"net"
"encoding/gob")

func sender(conn net.Conn){
	for{
		var msg string

		fmt.Scanf("%s",&msg)
	
		err := gob.NewEncoder(conn).Encode(msg)
	
		if err != nil{
			fmt.Println(err)
		}
	}
}

func handleRequest(conn net.Conn) {
  var msg string
  fmt.Println("handle start")
  for msg != "End"{
  err := gob.NewDecoder(conn).Decode(&msg)
  
  if err!= nil{
	fmt.Println(err)
	return
  }else{
  fmt.Println("Recieved", msg)}
  }
  
  conn.Close()
}

func main(){
	conn, _ := net.Dial("tcp", "127.0.0.1:3333")
	defer conn.Close()
	go handleRequest(conn)
	for{
		var msg string

		fmt.Scanf("%s",&msg)
	
		err := gob.NewEncoder(conn).Encode(msg)
	
		if err != nil{
			fmt.Println("2",err)
		}
	}
}
