package log

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "index_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	c := Config{}
	c.Segment.MaxIndexBytes = 1024
	idx, err := newIndex(f, c)
	require.NoError(t, err)
	_, _, err = idx.Read(-1)
	require.Error(t, err)
	require.Equal(t, f.Name(), idx.Name())
	entries := []struct {
		Off uint32
		Pos uint64
	}{
		{Off: 0, Pos: 0},
		{Off: 1, Pos: 10},
	}

	for _, want := range entries {
		t.Run(fmt.Sprintf("Off: %d,  Pos: %d", want.Off, want.Pos), func(t *testing.T) {
			err = idx.Write(want.Off, want.Pos)
			require.NoError(t, err)

			_, pos, err := idx.Read(int64(want.Off))
			require.NoError(t, err)
			require.Equal(t, want.Pos, pos)
		})
	}
}
