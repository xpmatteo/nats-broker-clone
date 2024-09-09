package nats

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestSession_PrintInfo(t *testing.T) {
	type fields struct {
		reader io.Reader
		writer io.Writer
		buf    []byte
	}
	tests := []struct {
		name           string
		fields         fields
		expectedErr    error
		expectedOutput string
	}{
		{
			name: "ok",
			fields: fields{
				writer: &bytes.Buffer{},
				reader: nil,
				buf:    make([]byte, 100),
			},
			expectedErr:    nil,
			expectedOutput: "INFO {\"maxPayload\": 100}\r\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewSession(test.fields.reader, test.fields.writer, 100)

			err := s.PrintInfo()

			if test.expectedErr == nil {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expectedOutput, test.fields.writer.(*bytes.Buffer).String())
		})
	}
}
