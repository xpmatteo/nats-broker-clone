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

		go serve(conn)
	}
}

const maxPayload = 100

// serve handles a single connection
func serve(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, maxPayload)
	for {
		err := interact(conn, buf)
		if err != nil {
			fmt.Println("Error from connection:", err)
			break
		}
	}
}

// interact handles a single interaction on a connection
func interact(conn io.ReadWriter, buf []byte) error {
	_, err := conn.Read(buf)
	if err != nil {
		return err
	}

	_, err = conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}
