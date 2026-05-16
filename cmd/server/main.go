package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	handleError(err)

	conn, err := listener.Accept()
	handleError(err)

	// we can print what we got
	readBuf := make([]byte, 1024*1024)
	n, err := conn.Read(readBuf)

	data := string(readBuf[0:n])
	fmt.Println(data)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err, "another one")
	}
}
