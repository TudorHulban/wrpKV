package kvbadger

import (
	"fmt"
	"sync"
	"testing"

	"github.com/TudorHulban/kv"
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

	kv0 := kv.KV{
		Key:   []byte(kPrefix + "x0"),
		Value: []byte("y0"),
	}
	assert.NoError(t,
		inMemoryStore.Set("", kv0),
	)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		kv1 := kv.KV{
			Key:   []byte(kPrefix + "x1"),
			Value: []byte("y1"),
		}

		assert.NoError(t,
			inMemoryStore.Set("", kv1),
		)

		wg.Done()
	}()

	go func() {
		kv2 := kv.KV{
			Key:   []byte(kPrefix + "x2"),
			Value: []byte("y2"),
		}

		assert.NoError(t,
			inMemoryStore.Set("", kv2),
		)

		wg.Done()
	}()

	wg.Wait()

	reconstructedExistingPrefix, errGet := inMemoryStore.GetKVByPrefix(
		"",
		[]byte(kPrefix),
	)
	assert.NoError(t, errGet)
	assert.Len(t, reconstructedExistingPrefix, 3)
	assert.EqualValues(t,
		reconstructedExistingPrefix[0].Key,
		kv0.Key,

		fmt.Sprint(reconstructedExistingPrefix[0]),
	)

	reconstructedNonExistingPrefix, errBadPrefix := inMemoryStore.GetKVByPrefix(
		"",
		[]byte("xxxxxxxxxx"),
	)
	assert.NoError(t, errBadPrefix)
	assert.Zero(t, reconstructedNonExistingPrefix)
}
