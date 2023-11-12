package kvbadger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	inMemoryStore, err := NewBStoreInMemoryNoLogging()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, inMemoryStore.Close())
	}()

	key := []byte("x")
	value := []byte("y")

	assert.NoError(t,
		inMemoryStore.Set(key, value),
	)

	value0, errGet0 := inMemoryStore.GetValueFor(key)
	assert.NoError(t, errGet0)
	assert.Equal(t, value0, value)

	assert.NoError(t,
		inMemoryStore.DeleteKVBy(key),
	)

	value1, errGet1 := inMemoryStore.GetValueFor(key)
	assert.Error(t, errGet1)
	assert.Zero(t, value1)
}
