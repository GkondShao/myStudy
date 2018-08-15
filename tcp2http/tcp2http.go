package main


import(
	"log"
	"net"
	"os"
	"myStudy/tcp2http/util"
	"myStudy/tcp2http/constant"
)




type  httpHeader struct{
	Version string

}


func checkError(err error){
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
}





func handlerConnection(conn net.Conn){
	log.Printf("recieve from :%s,type :%s\n",conn.RemoteAddr().String(),conn.RemoteAddr().Network())

	// 获取http 信息
	buf,err  := util.ReadAllFromConn(conn)
	if err != nil{
		conn.Write(ResponseWithCode(constant.StatusBadRequest))
	}
	
}


func main(){
	tcpListener,err := net.Listen("tcp",":8081")
	checkError(err)

	defer tcpListener.Close()

	log.Println("start listen server...")

	for{
		conn,err := tcpListener.Accept()
		if err != nil{
			log.Printf("listen error :%v",err)
			continue
		}

		defer conn.Close()
		handlerConnection(conn)
	}

}