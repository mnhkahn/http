package main

import (
	"flag"
	"fmt"
	"net"
)

var portFlag = flag.Int("p", 8080, "Http Server Port")

func main() {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "", *portFlag))
	if err != nil {
		panic(err)
	}

	fmt.Printf("<<<Server Accepting on Port %d>>>\n\n", *portFlag)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("<<<Request From %s>>>\n", conn.RemoteAddr())
	conn.Write([]byte("HTTP/1.1 200 OK\nDate: Thu, 10 Dec 2015 06:29:08 GMT\nContent-Type: text/html; charset=utf-8\nServer: Cyeam\nContent-length:12\n\nhello, world\r\n\r\n"))
}
