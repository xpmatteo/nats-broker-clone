package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
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

		go func() {
			defer conn.Close()
			serve(conn)
		}()
	}
}

func serve(conn io.ReadWriter) {
	buf := make([]byte, 10)
	for {
		nr, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		command := strings.TrimSpace(string(buf[:nr]))
		response := interact(command)

		_, err = conn.Write([]byte(response + "\r\n"))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}

func interact(command string) string {
	response := command
	return response
}
