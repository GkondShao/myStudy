package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func sender(tcpAddr *net.TCPAddr, i int, count chan int) {
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		log.Fatal("No.%d send failed:", err)
	}
	conn.Write([]byte(fmt.Sprintf("send message no %d", i)))
	log.Printf("Send %d", i)
	count <- 1
}

func main() {

	count := make(chan int, 10)
	server := "127.0.0.1:8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error :%s", err.Error())
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		go sender(tcpAddr, i, count)
	}

	time.Sleep(10 * time.Second)

	for j := 1; j < 10; j++ {
		<-count
	}
}
