package main

import (
	"Socket/SocketB/protocol"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func send(conn net.Conn) {
	for i := 0; i < 100; i++ {
		session := getSession()
		words := "{\"ID\":" + "\"" + strconv.Itoa(i) + "\",\"Session\":" + "\"" + session + "2015073109532345\",\"Meta\":" + "Gkond" + "\"Content\":" + "\"" + fmt.Sprintf("message %d", i) + "\"" + "}"
		//fmt.Println(words)
		conn.Write(protocol.Enpack([]byte(words)))
	}
	fmt.Println("send over")

	defer conn.Close()
}

func getSession() string {
	gs1 := time.Now().Unix()
	gs2 := strconv.FormatInt(gs1, 10)
	return gs2
}

func main() {
	server := "localhost:8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)

	if err != nil {
		log.Fatal("Address resolve error : ", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		log.Fatal("server connect error : ", err)
	}

	log.Println("connect Success!")

	send(conn)

}
