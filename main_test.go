package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

type failWriter struct {
}

var writeFailed = fmt.Errorf("write failed")

func (f *failWriter) Write(p []byte) (n int, err error) {
	return 0, writeFailed
}

func (f *failWriter) Read(p []byte) (n int, err error) {
	return 0, nil
}

type bufReadWriter struct {
	buf []byte
}

func (b *bufReadWriter) Write(p []byte) (n int, err error) {
	nw := copy(b.buf, p)
	return nw, nil
}

func (b *bufReadWriter) Read(p []byte) (n int, err error) {
	nr := copy(p, b.buf)
	return nr, nil
}

type bufWriter struct {
	buf []byte
	nw  int
}

func (b *bufWriter) Write(p []byte) (n int, err error) {
	b.buf = make([]byte, len(p))
	nw := copy(b.buf, p)
	return nw, nil
}

func (b *bufWriter) Read(p []byte) (n int, err error) {
	panic("should not be called")
}

func Test_printInfo_returns_error_on_write_error(t *testing.T) {
	failWriter := &failWriter{}
	err := printInfo(failWriter, make([]byte, maxPayload))
	assert.ErrorIs(t, err, writeFailed)
}

func Test_printInfo_starts_with_info_declaration(t *testing.T) {
	writer := &bufWriter{}

	err := printInfo(writer, make([]byte, maxPayload))

	expected := "INFO {}\r\n"
	assert.NoError(t, err)
	assert.Equal(t, []byte(expected), writer.buf)
}

func Test_expectConnection(t *testing.T) {
	tests := []struct {
		name           string
		readWriter     io.ReadWriter
		expectedError  error
		expectedOutput string
	}{
		{"ok", send("CONNECT {}\r\n"), nil, "+OK\r\n"},
		//{"bad syntax", "FOO\r\n"},
		//{"timeout", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := expectConnection(test.readWriter)

			if test.expectedError == nil && err != nil {
				assert.NoError(t, err)
				return
			}
			if test.expectedError != nil && err == nil {
				assert.ErrorIs(t, err, test.expectedError)
				return
			}
			expect(t, test.readWriter, test.expectedOutput)
		})
	}
}

func expect(t *testing.T, rw io.ReadWriter, expectedOutput string) {
	buf := make([]byte, maxPayload)

	n, err := rw.Read(buf)

	assert.NoError(t, err)
	assert.Equal(t, []byte(expectedOutput), buf[:n])
}

func send(input string) io.ReadWriter {
	rw := &bufReadWriter{}
	_, _ = rw.Write([]byte(input))
	return rw
}
