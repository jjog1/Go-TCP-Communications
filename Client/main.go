package main

import ( "fmt"
"net"
"encoding/gob")

type Person struct{
	Name string
	Age int
}


func handleRequest(conn net.Conn) {
  var msg string
  fmt.Println("handles start")
  dec:=gob.NewDecoder(conn)
  for msg != "End"{
  err := dec.Decode(&msg)
  
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

		length,_ := fmt.Scanf("%s",&msg)

		fmt.Println("Length: ",length)
		person:= Person{msg, 20}
		
		fmt.Println("Client",person)
		
		//Scanf seems to be reading a second blank line
		//for every input. 
		if length > 0{
			// Creating a new encoder doesn't seem the best 
			// way to do this. 
			err := gob.NewEncoder(conn).Encode(&person)
				
			if err != nil{
				fmt.Println("2",err)
			}
			
			length = 0
		}
	}
}
