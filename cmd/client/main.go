package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

//small implementation
// a program that serves a file over a network connection.
// Two versions. You'll feel the difference.

// sends data over a connection
func main() {
	// returns a pointer to the Go os.File struct in memory.
	// The os.File struct wraps the kernel's file
	dataFile, err := os.Open("data.txt")
	handleError(err)

	//lets create a connection next
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	handleError(err)

	// copy way - benchmark
	startCopy := time.Now()
	data := make([]byte, 1024*1024)
	n, err := dataFile.Read(data)
	handleError(err)
	conn.Write(data[0:n])
	elapsedCopy := time.Since(startCopy)
	fmt.Printf("The copy way took %v to run.\n\n", elapsedCopy)

	// zero copy way - benchmark
	startZeroCopy := time.Now()
	tcpConn := conn.(*net.TCPConn) // type Assertion to get tcpconn
	socketFile, err := tcpConn.File()
	handleError(err)

	src := dataFile.Fd()
	dst := socketFile.Fd()
	offset := int64(0) // where to start sending from the file
	fileinfo, err := dataFile.Stat()
	handleError(err)
	count := fileinfo.Size() // how many bytes to send
	syscall.Sendfile(int(dst), int(src), &offset, int(count))
	elapsedZeroCopy := time.Since(startZeroCopy)
	fmt.Printf("The zero copy way took %v to run.\n\n", elapsedZeroCopy)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
