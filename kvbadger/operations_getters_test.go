package kvbadger

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByPrefix(t *testing.T) {
	inMemoryStore, err := NewBStoreInMemoryNoLogging()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t,
			inMemoryStore.Close(),
		)
	}()

	kPrefix := "prefix-"

	kv0 := KV{
		Key:   []byte(kPrefix + "x0"),
		Value: []byte("y0"),
	}
	assert.NoError(t,
		inMemoryStore.SetKV(kv0),
	)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		kv1 := KV{
			Key:   []byte(kPrefix + "x1"),
			Value: []byte("y1"),
		}

		assert.NoError(t,
			inMemoryStore.SetKV(kv1),
		)

		wg.Done()
	}()

	go func() {
		kv2 := KV{
			Key:   []byte(kPrefix + "x2"),
			Value: []byte("y2"),
		}

		assert.NoError(t,
			inMemoryStore.SetKV(kv2),
		)

		wg.Done()
	}()

	wg.Wait()

	fetchedKeyValues, errGet := inMemoryStore.GetKVByPrefix([]byte(kPrefix))
	assert.NoError(t, errGet)
	assert.Equal(t, len(fetchedKeyValues), 3)
	assert.Contains(t, fetchedKeyValues, kv0)

	badPrefix, errBadPrefix := inMemoryStore.GetKVByPrefix([]byte("xxxxxxxxxx"))
	assert.NoError(t, errBadPrefix)
	assert.Equal(t, 0, len(badPrefix))
}
