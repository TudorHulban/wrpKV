package kvbadger

import (
	"strconv"
	"testing"
	"time"

	"github.com/TudorHulban/kv"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	inMemoryStore, err := NewBStoreInMemoryNoLogging()
	assert.NoError(t, err)

	defer func() {
		assert.NoError(t,
			inMemoryStore.Close(),
		)
	}()

	key := []byte("prefix-" + "x")
	value := []byte("y")

	assert.NoError(t, inMemoryStore.Set(kv.KV{
		Key:   key,
		Value: value,
	}))

	updateValue := []byte("z")
	assert.NoError(t,
		inMemoryStore.Set(kv.KV{
			Key:   key,
			Value: updateValue,
		}),
	)

	reconstructedValue, errGet := inMemoryStore.GetValueFor(key)
	assert.NoError(t, errGet)

	t.Logf("value: %s", updateValue)
	t.Logf("fetched: %s", reconstructedValue)
	assert.Equal(t, updateValue, reconstructedValue)
}

func TestOnClosedStore(t *testing.T) {
	inMemoryStore, err := NewBStoreInMemoryNoLogging()
	assert.NoError(t, err)
	assert.NoError(t, inMemoryStore.Close())

	kv := kv.KV{
		Key:   []byte("x"),
		Value: []byte("y"),
	}

	assert.Error(t,
		inMemoryStore.Set(kv),
	)
}

func TestTTL(t *testing.T) {
	inMemoryStore, err := NewBStoreInMemoryNoLogging()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, inMemoryStore.Close())
	}()

	key := []byte("prefix-" + "x")
	value := []byte("y")

	ttlSeconds := 1

	assert.NoError(t,
		inMemoryStore.SetTTL(
			kv.KV{
				Key:   key,
				Value: value,
			},
			uint(ttlSeconds),
		),
	)

	time.Sleep(time.Duration(ttlSeconds+1) * time.Second)

	_, errGet := inMemoryStore.GetValueFor(key)
	assert.Error(t, errGet)
}

// cpu: AMD Ryzen 7 5800H with Radeon Graphics, Ubuntu / Debian
// BenchmarkSet-16           112372             10600 ns/op            1449 B/op         31 allocs/op
// BenchmarkSetKV-16         120267             10219 ns/op            1451 B/op         31 allocs/op

// cpu: AMD Ryzen 7 5800H with Radeon Graphics, Fedora
// BenchmarkSet-16            26874             44579 ns/op            1482 B/op         31 allocs/op
// BenchmarkSetKV-16          26198             45505 ns/op            1490 B/op         31 allocs/op
func BenchmarkSet(b *testing.B) {
	inMemoryStore, errCr := NewBStoreInMemoryNoLogging()
	if errCr != nil {
		b.FailNow()
	}
	defer inMemoryStore.Close()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		inMemoryStore.Set(
			kv.KV{
				Key:   []byte(strconv.Itoa(i)),
				Value: []byte("x"),
			},
		)
	}
}

func BenchmarkSetKV(b *testing.B) {
	inMemoryStore, errCr := NewBStoreInMemoryNoLogging()
	if errCr != nil {
		b.FailNow()
	}
	defer inMemoryStore.Close()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		kv := kv.KV{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("x"),
		}

		inMemoryStore.Set(kv)
	}
}
