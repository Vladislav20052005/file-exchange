package recieve

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func Recieve(name string, addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Printf("Programm is waiting for file %s...\n", name)

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for {
		bufferSize := 1024 * 1024
		buffer := make([]byte, bufferSize)
		n, err := conn.Read(buffer)

		if n > 0 {
			_, err := writer.Write(buffer[:n])
			if err != nil {
				panic(err)
			}
			err = writer.Flush()
			if err != nil {
				panic(err)
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}

		}

	}

	fmt.Println("Successfully recieved file.")
}
