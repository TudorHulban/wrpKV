package kvnuts

import (
	"testing"

	"github.com/TudorHulban/kv/definition"
	"github.com/stretchr/testify/require"
)

func TestGetByPrefix(t *testing.T) {
	var store definition.KVStore

	var errCr error

	store, errCr = NewStoreInMemory(_segmentSizeTests)
	require.NoError(t, errCr)
	require.NotZero(t, store)
}
