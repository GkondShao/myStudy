package main

import (
	SocketUtil "Socket/SocketUtil"
	"fmt"
	"log"
	"net"
	"os"
)

func checkErr(err error) {

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}

func Log(v ...interface{}) {
	log.Println(v)
}

func handleConnection(conn net.Conn) {

	buf, err := SocketUtil.ReadAll(conn)
	if err != nil {
		log.Fatal(conn.RemoteAddr().String(), " connection error : ", err)
	}
	Log(conn.RemoteAddr().String(), "receive data string:", string(buf))

}

func main() {

	listen, err := net.Listen("tcp", ":8080")

	checkErr(err)
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " tcp connection success!")
		go handleConnection(conn)

	}

}
