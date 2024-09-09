package nats

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func Test_PrintInfo_ok(t *testing.T) {
	tests := []struct {
		name           string
		session        *Session
		expectedOutput string
	}{
		{
			name:           "ok 100",
			session:        NewSession(nil, &bytes.Buffer{}, 100),
			expectedOutput: "INFO {\"maxPayload\": 100}\r\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.session.PrintInfo()

			assert.NoError(t, err)
			assert.Equal(t, test.expectedOutput, test.session.writer.(*bytes.Buffer).String())
		})
	}
}

func Test_PrintInfo_error(t *testing.T) {
	s := NewSession(nil, returningError(io.ErrShortWrite), 100)

	err := s.PrintInfo()

	assert.ErrorIs(t, err, io.ErrShortWrite)
}

type WriterFunc func(p []byte) (n int, err error)

func (f WriterFunc) Write(p []byte) (n int, err error) {
	return f(p)
}

func returningError(err error) io.Writer {
	return WriterFunc(func(p []byte) (n int, e error) {
		return 0, err
	})
}
