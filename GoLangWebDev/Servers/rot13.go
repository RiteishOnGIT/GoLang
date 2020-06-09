package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main(){
	li, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil{
			log.Fatalln(err)
			continue
		}
		go handle(conn)
	}
}


func handle(conn net.Conn){
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		ln := strings.ToLower(scan.Text())
		bs := []byte(ln)
		r := rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
	}
}

func rot13( s []byte)[]byte{
	var r13 = make([]byte , len(s))
	for i, v := range s {
		// ascii 97 - 122
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13

}