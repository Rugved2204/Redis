package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {    
    listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Printlnt("Error starting the server:", err)

	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Printl("Error Accepting connection:", err)
	}

	for {
		buf := make([]byte, 1024)


		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("error reading from connection:", err)
			}
		}

		conn.Write([]byte("+OK"))
	}
}
