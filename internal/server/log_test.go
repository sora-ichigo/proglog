package server

import (
	"testing"
)

func TestNewLog(t *testing.T) {
	NewLog()
}

func TestAppend(t *testing.T) {
	log := NewLog()
	r := Record{
		Value: []byte{},
	}

	_, err := log.Append(r)
	if err == nil {
		t.Fatalf("failed to append %v", err)
	}
}
