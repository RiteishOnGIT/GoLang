package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main(){
	li ,err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalln(err)
	}
	defer li.Close()
	for{
		conn ,err := li.Accept()
		if err != nil{
			log.Fatalln(err)
		}
		go handle(conn)
	}

}

func handle(con net.Conn){
	scanner := bufio.NewScanner(con)
	for scanner.Scan(){
		li := scanner.Text()
		fmt.Println(li)
	}
	defer con.Close()
	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}