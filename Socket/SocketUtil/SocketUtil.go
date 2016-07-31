package Socket

import (
	"io"
	"net"
)

func ReadAll(conn net.Conn) ([]byte, error) {
	tmpBuf := make([]byte, 1024)
	Buf := make([]byte, 0)

	for {
		n, err := conn.Read(tmpBuf)

		if err != nil {
			if err != io.EOF {
				return nil, err
			}
		}

		Buf = append(Buf, tmpBuf[:n]...)

		if err == io.EOF {
			return Buf, err

		}
	}

}
