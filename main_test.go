package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func Test_serve_returns_error_on_write_error(t *testing.T) {
	failWriter := &failWriter{}
	err := interact(failWriter, make([]byte, maxPayload))
	assert.ErrorIs(t, err, writeFailed)
}
