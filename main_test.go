package main

import (
	"fmt"
	"io"
	"testing"
)

//func Test_serve(t *testing.T) {
//	type args struct {
//		conn io.ReadWriter
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			serve(test.args.conn)
//		})
//	}
//}

type failWriter struct {
}

func (f *failWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("write failed")
}

func (f *failWriter) Read(p []byte) (n int, err error) {
	return 0, nil
}

func test_serve_returns_error_on_read_error(t *testing.T) {
	failWriter := &failWriter{}
	var conn io.ReadWriter = failWriter
	buf := make([]byte, maxPayload)
	for {
		if err := interact(conn, buf); err != nil {
			fmt.Println("Error from connection:", err)
			break
		}
	}
}
