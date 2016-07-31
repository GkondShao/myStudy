package protocol

import (
	"bytes"
	"encoding/binary"
	"log"
)

const (
	ConstHeader       = "Headers"
	ConstHeaderLength = 7

	//int2byte  length = 4
	ConstMLength = 4
)

func Byte2Int(buffer []byte) int {
	bytesBuffer := bytes.NewBuffer(buffer)

	var x int32

	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

func Int2Byte(n int) []byte {
	x := int32(n)
	byteBuffer := bytes.NewBuffer([]byte{})

	binary.Write(byteBuffer, binary.BigEndian, x)
	return byteBuffer.Bytes()
}

func Enpack(message []byte) []byte {
	bytes := append(append([]byte(ConstHeader), Int2Byte(len(message))...), message...)
	log.Println(string(bytes))
	return bytes
}

func Depack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int

	for i = 0; i < length; i++ {
		//buffer[ConstHeaderLength+ConstMLength：] is the message
		if length < i+ConstHeaderLength+ConstMLength {
			break
		}

		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			messageLength := Byte2Int(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstMLength])
			//log.Println(messageLength)
			//wait if the message is uncomplete
			if length < i+ConstHeaderLength+ConstMLength+messageLength {
				break
			}

			data := buffer[i+ConstHeaderLength+ConstMLength : i+ConstHeaderLength+ConstMLength+messageLength]
			//log.Println(string(data))
			readerChannel <- data
		}
	}

	if i == length {

		return make([]byte, 0)
	}
	//log.Println(string(buffer[i:]))
	return buffer[i:]
}
