package main

import (
	"fmt"
	"os"

	recieve "file_exchange/recieve"
	send "file_exchange/send"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("send ip-address:port path \nrecieve :port path")
		return
	}
	var run_type string = os.Args[1]

	switch run_type {
	case "send":
		var (
			addr string = os.Args[2]
			path string = os.Args[3]
		)
		send.Send(path, addr)
	case "recieve":
		var (
			port string = os.Args[2]
			path string = os.Args[3]
		)
		recieve.Recieve(path, port)
	default:
		fmt.Println("Unidentified argument: ", run_type)
		fmt.Println("Try recieve or send")
		return
	}
}
