// TaskServer project main.go
package main

import (
	. "TaskServer/Servers"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	manager := new(TaskManager)
	rpc.Register(manager)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go jsonrpc.ServeConn(conn)

	}
}
