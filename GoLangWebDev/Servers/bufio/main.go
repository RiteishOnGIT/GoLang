package main


import (
	"bufio"
	"fmt"
	"log"
	"net"
)


func main(){
	li, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil{
			log.Fatalln(err)
		}
		go handle (conn)
	}
}

func handle(conn net.Conn){
	Scanner := bufio.NewScanner(conn)
	for Scanner.Scan(){
		text  := Scanner.Text()
		fmt.Print(text)
	}
}