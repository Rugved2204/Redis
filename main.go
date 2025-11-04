package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"	
	"log"
)

const (
	port = ":6379"
)

func main() {
	fmt.Printf("My redis server is starting on port: %s\n", port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening for connections")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}

		fmt.Printf("New Client is connected: %s\n", conn.RemoteAddr())

		go handleclient(conn)
	}


}

func handleclient(conn net.Conn) {
	defer func(){
		fmt.Printf("Client disconnected: %s\n", conn.RemoteAddr())
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				return
			}
			log.Printf("Error reading the message: %s", err)
			return
		}

		message = strings.TrimSpace((message))

		if message == "" {
			continue
		}

		fmt.Printf("Received from %s: %s\n", conn.RemoteAddr(), message)

		response := fmt.Sprintf("ECHO: %s", message)
		_, err = writer.WriteString(response)
		if err != nil {
			log.Println("Failed to tell the response")
		}

		writer.Flush()
	}
}
