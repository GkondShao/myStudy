package main

import (
	SocketUtil "Socket/SocketUtil"
	"fmt"
	"io"
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
	if err != nil && err != io.EOF {
		log.Fatal(conn.RemoteAddr().String(), " connection error : ", err)
	}
	//test the return
	log.Println(err)
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
