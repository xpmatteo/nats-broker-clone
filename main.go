package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	const port = "4444"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port " + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	_, err := conn.Write([]byte("HELLO\r\n"))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
	}

	// Close the connection
	_ = conn.Close()
}
