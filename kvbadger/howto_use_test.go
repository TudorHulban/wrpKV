package kvbadger_test

import (
	"testing"

	"github.com/TudorHulban/kv/kvbadger"
	"github.com/stretchr/testify/require"
)

func TestHowToUse(t *testing.T) {
	inMemoryStore, err := kvbadger.NewBStoreInMemoryNoLogging()
	require.NoError(t, err)

	key := []byte("x")
	value := []byte("y")

	require.NoError(t, inMemoryStore.Set(key, value))

	fetchedValue, errGet := inMemoryStore.GetValueFor(key)
	require.NoError(t, errGet)
	require.Equal(t, fetchedValue, []byte(value))
}
