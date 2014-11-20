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
	
		person:= Person{msg, 20}
		
		fmt.Println(person)
		
		err := gob.NewEncoder(conn).Encode(person)
		//err := gob.NewEncoder(conn).Encode(msg)
	
		if err != nil{
			fmt.Println("2",err)
		}
	}
}
