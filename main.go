package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":"+"4444")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Send the string "HELLO\r\n" to the client
		_, err = conn.Write([]byte("HELLO\r\n"))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
		}

		// Close the connection
		conn.Close()
	}
}
