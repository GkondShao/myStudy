package main

import (
	"Socket/SocketB/protocol"
	"io"
	"log"
	"net"
)

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			log.Println(string(data))
		}
	}
}

func handleConnection(conn net.Conn) {
	//save the partly message
	tmpBuffer := make([]byte, 0)

	//dual 16 message mostly
	readerChannel := make(chan []byte, 16)

	go reader(readerChannel)

	buffer := make([]byte, 10)

	for {
		n, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			log.Println(conn.RemoteAddr().String(), " connect error : ", err)

		}

		tmpBuffer = protocol.Depack(append(tmpBuffer, buffer[:n]...), readerChannel)
		//log.Println(string(tmpBuffer))
	}

	defer conn.Close()

}

func main() {
	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("listen error : ", err.Error())
	}

	defer listen.Close()

	log.Println("waiting for connnect")

	for {
		conn, err := listen.Accept()

		if err != nil {
			continue
		}

		log.Println(conn.RemoteAddr().String(), " connected")

		go handleConnection(conn)

	}
}
