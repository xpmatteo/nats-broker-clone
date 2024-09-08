package main

import (
	"fmt"
	"io"
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

		go func() {
			defer conn.Close()
			serve(conn)
		}()
	}
}

func serve(conn io.ReadWriter) {
	buf := make([]byte, 10)
	nr, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}
	logBuffer(buf, nr)

	_, err = conn.Write(buf)
	if err != nil {
		fmt.Println("Error writing to connection:", err)
	}
}

// logBuffer prints the characters in buf one by one
func logBuffer(buf []byte, nr int) {
	for i := range nr {
		ch := buf[i]
		fmt.Printf("%02X: %c\n", ch, ch)
	}
}
