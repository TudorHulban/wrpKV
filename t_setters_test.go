package badgerwrap

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/TudorHulban/log"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	l := log.New(log.DEBUG, os.Stderr, true)

	inmemStore, err := NewBStoreInMem(l)
	assert.Nil(t, err)
	defer func() {
		assert.Nil(t, inmemStore.Close())
	}()

	kPrefix := "prefix-"
	kv := KV{[]byte(kPrefix + "x"), []byte("y")}

	// test insert
	assert.Nil(t, inmemStore.Set(kv))

	// test update
	kv.value = []byte("z")
	assert.Nil(t, inmemStore.Set(kv))

	v, errGet := inmemStore.GetVByK(kv.key)
	assert.Nil(t, errGet)
	assert.Equal(t, v, []byte(kv.value))
}

func TestClose(t *testing.T) {
	inmemStore, err := NewBStoreInMemNoLogging()
	assert.Nil(t, err)
	assert.Nil(t, inmemStore.Close())

	// test insert on closed store.
	kv := KV{[]byte("x"), []byte("y")}
	errSet := inmemStore.Set(kv)
	assert.Error(t, errSet)
}

func TestTTL(t *testing.T) {
	l := log.New(log.DEBUG, os.Stderr, true)

	inmemStore, err := NewBStoreInMem(l)
	assert.Nil(t, err)
	defer func() {
		assert.Nil(t, inmemStore.Close())
	}()

	kPrefix := "prefix-"
	kv := KV{[]byte(kPrefix + "x"), []byte("y")}
	ttl := 1

	errSet := inmemStore.SetTTL(kv, uint(ttl))
	assert.Nil(t, errSet)

	time.Sleep(time.Duration(ttl+1) * time.Second)
	_, errGet := inmemStore.GetVByK(kv.key)
	assert.Error(t, errGet)
}

// BenchmarkSet-4   	   19934	     59591 ns/op	    1367 B/op	      34 allocs/op
func BenchmarkSet(b *testing.B) {
	inmemStore, _ := NewBStoreInMemNoLogging()
	defer func() {
		inmemStore.Close()
	}()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		inmemStore.Set(KV{
			[]byte(strconv.Itoa(i)),
			[]byte("x"),
		})
	}
}
