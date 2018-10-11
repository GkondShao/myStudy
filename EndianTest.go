package main

import(
	"log"
	"unsafe"
)

func main(){
	a :=0x1234

	b :=  *(*int8)(unsafe.Pointer(&a))

	if b == 0x12{
		log.Println("Big Endian")
	}else if b == 0x34{
		log.Println("little Endian")
	}else{
		log.Println("error")
	}
}