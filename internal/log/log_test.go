package log

import (
	"os"
	"testing"

	log_v1 "github.com/igsr5/proglog/api/v1"
	"github.com/stretchr/testify/require"
)

func TestNewLog(t *testing.T) {
	testNewLog(t)
}

func TestLogAppend(t *testing.T) {
	l := testNewLog(t)
	r := &log_v1.Record{Value: []byte("hoge")}
	testLogAppend(t, l, r)
}

func TestLogRead(t *testing.T) {
	l := testNewLog(t)
	r := &log_v1.Record{Value: []byte("hoge")}
	off := testLogAppend(t, l, r)
	got, err := l.Read(off)
	require.Equal(t, got.Value, r.Value)
	require.NoError(t, err)
}

func testNewLog(t *testing.T) *Log {
	t.Helper()

	dir := os.TempDir()
	c := Config{}
	l, err := NewLog(dir, c)
	require.NoError(t, err)
	err = l.setup()
	require.NoError(t, err)

	return l
}

func testLogAppend(t *testing.T, l *Log, r *log_v1.Record) uint64 {
	off, err := l.Append(r)
	require.NoError(t, err)

	return off
}
