package badgerwrap

import (
	"strconv"
	"testing"
	"time"

	"github.com/TudorHulban/loginfo"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	l, errLog := loginfo.New(2)
	assert.Nil(t, errLog)

	inmemStore, err := NewBStoreInMem(&l)
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
	l, errLog := loginfo.New(2)
	assert.Nil(t, errLog)

	inmemStore, err := NewBStoreInMem(l)
	assert.Nil(t, err)
	assert.Nil(t, inmemStore.Close())

	// test insert on closed store.
	kv := KV{[]byte("x"), []byte("y")}
	errSet := inmemStore.Set(kv)
	assert.Error(t, errSet)
}

func Test3TTL(t *testing.T) {
	l, errLog := loginfo.New(2)
	assert.Nil(t, errLog)

	inmemStore, err := NewBStoreInMem(l)
	assert.Nil(t, err)
	defer func() {
		assert.Nil(t, inmemStore.Close())
	}()

	kPrefix := "prefix-"
	kv := KV{[]byte(kPrefix + "x"), []byte("y")}
	ttl := 1

	errSet := inmemStore.SetTTL(kv, uint8(ttl))
	assert.Nil(t, errSet)

	time.Sleep(time.Duration(ttl+1) * time.Second)
	_, errGet := inmemStore.GetVByK(kv.key)
	assert.Error(t, errGet)
}

// BenchmarkSet-4   	   11239	     91726 ns/op	   22480 B/op	      59 allocs/op
func BenchmarkSet(b *testing.B) {
	inmemStore, _ := NewBStoreInMemNoLogging()
	defer func() {
		inmemStore.Close()
	}()

	for i := 0; i < b.N; i++ {
		inmemStore.Set(KV{
			[]byte(strconv.Itoa(i)),
			[]byte("x"),
		})
	}
}
