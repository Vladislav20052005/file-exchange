package send

import (
	"fmt"
	"io"
	"net"
	"os"
)

func Send(path string, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bufferSize := 1024 * 1024
	buffer := make([]byte, bufferSize)

	for {
		n, err := file.Read(buffer)

		if n > 0 {
			_, err = conn.Write(buffer[:n])
			if err != nil {
				panic(err)
			}
		}

		if err != nil {
			if err == io.EOF {
				fmt.Println("End of file reached.")
				break
			}
			panic(err)
		}
	}
}
