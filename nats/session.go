package nats

import (
	"fmt"
	"io"
)

type Session struct {
	reader io.Reader
	writer io.Writer
	buf    []byte
}

func NewSession(reader io.Reader, writer io.Writer, maxPayload int) *Session {
	return &Session{
		reader: reader,
		writer: writer,
		buf:    make([]byte, maxPayload),
	}
}

func (s *Session) PrintInfo() error {
	_, err := s.writer.Write([]byte(fmt.Sprintf("INFO {\"maxPayload\": %d}\r\n", len(s.buf))))
	return err
}
