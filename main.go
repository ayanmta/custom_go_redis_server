package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("connection started on port :6379")
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("error connecting server")
		return
	}
	conn, _ := l.Accept()
	defer conn.Close()

	for {
		buff := make([]byte, 1024)
		_, err := conn.Read(buff)
		if err != nil {
			fmt.Println("error establishing connection and reading from client ", err.Error())
			os.Exit(1)
		}
		conn.Write([]byte("+lifeOK\r\n"))
	}
}
