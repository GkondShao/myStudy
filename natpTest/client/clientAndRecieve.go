package main

import(
	"net"
	"log"
	"time"
)


const NAME = "CLIENT1"
var  rip = "xxx.xxx.xxx.xxx:8888"

func main(){

	log.SetFlags(log.Ldate|log.Llongfile)

	// tip anyone
	port := ":8081"

	udpaddr,err := net.ResolveUDPAddr("udp", port)
	if err != nil{
		log.Fatalln("resolve laddr fail:",err)
	}

	log.Println(udpaddr)

	conn,err := net.ListenUDP("udp", udpaddr)
	if err != nil{
		log.Fatalf("listen port[%d] fail:%v", port,err)
	}

	defer conn.Close()

	// recv notify
	go func(){
		var recieve []byte

		for{
			recieve = make([]byte,1024)
			n,raddr,err := conn.ReadFromUDP(recieve)
			if err != nil{
				log.Printf("read from udp err: %v",err)
				continue
			}
			log.Printf("recv notify from[%s:%d] %d byte:%s", raddr.IP,raddr.Port,n,string(recieve[:n]))
		}
	}()


	// heart
	go func(){
		raddr,err := net.ResolveUDPAddr("udp", rip)
		if err != nil{
			log.Fatalf("resolve raddr fail:%v", err)
		}

		for{
			_,err := conn.WriteToUDP([]byte("CLIENT1"), raddr)
			if err != nil{
				log.Printf("heart to %s err:%v\n",rip,err)
				continue
			}

			time.Sleep(2*time.Second)
		}
	}()

	

	for{
		log.Println("running...")
		time.Sleep(50*time.Second)
	}

}