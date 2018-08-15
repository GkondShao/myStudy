package util



import(
	"net"
	"io"
)


/**
ReadAllFromConn  read from connection  
return slice when success,otherwise return nil and error
*/

func ReadAllFromConn(conn net.Conn)([]byte ,error){

	buf := make([]byte,0,2048)
	len := 0

	for{
		n,err := conn.Read(buf)
		if n > 0{
			len += n
		}

		if err != nil{
			if err != io.EOF{
				return nil,err
			}
		}


		return buf[:len],nil

	}
}


func ResponseWithCode(code int){
	
}