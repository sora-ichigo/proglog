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
	if err != nil {
		t.Fatalf("failed to append %v", err)
	}
}

func TestRead(t *testing.T) {
	log := NewLog()
	r := Record{
		Value: []byte{},
	}

	offset, err := log.Append(r)
	if err != nil {
		t.Fatalf("failed to append %v", err)
	}

	r, err = log.Read(offset)
	if err != nil {
		t.Fatalf("failed to read. err: %v", err)
	}
}
