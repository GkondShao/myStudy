package main

import(
	"net"
	"log"
	"time"
	"fmt"
)


var clients map[string]string

func main(){
	clients = make(map[string]string)
	port := ":8888"


	udpaddr,err := net.ResolveUDPAddr("udp", port)
	if err != nil{
		log.Fatalln("resolve laddr fail:",err)
	}


	conn,err := net.ListenUDP("udp", udpaddr)
	if err != nil{
		log.Fatalf("listen port[%d] fail:%v", port,err)
	}

	defer conn.Close()


	// recv heart
	go func(){
		
		var recieve []byte

		for{
			recieve = make([]byte,1024)
			n,raddr,err := conn.ReadFromUDP(recieve)
			if err != nil{
				log.Printf("read from udp err: %v",err)
				continue
			}

			log.Printf("recv heart from[%s:%d] %d byte:%s", raddr.IP,raddr.Port,n,string(recieve[:n]))
			clients[string(recieve[:n])] = fmt.Sprintf("%s%s%d",raddr.IP,":",raddr.Port)
		}
	}()

	// notify
	go func(){
		for{
			for k,v := range clients{
				raddr,err := net.ResolveUDPAddr("udp", v)
				if err != nil{
					log.Fatalf("get raddr[%s] fail:%v", k,err)
					continue
				}
				conn.WriteToUDP([]byte("notify"),raddr)
				log.Println("send to client:",k,v)
			}

			time.Sleep(2*time.Second)
		}
	}()

	for{
		log.Println("listening...")
		time.Sleep(50*time.Second)
	}

}