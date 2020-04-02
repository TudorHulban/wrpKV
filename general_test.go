package badgerwrap

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test1Set(t *testing.T) {
	inmemStore, err := NewBStore("", false)
	defer inmemStore.Close()

	a := assert.New(t)
	a.Nil(err)

	kv := KV{"prefix", "x", "y"}

	// test insert
	errSet := inmemStore.Set(kv)
	a.Nil(errSet)

	// test update
	kv.value = "z"
	errUpdate := inmemStore.Set(kv)
	a.Nil(errUpdate)

	v, errGet := inmemStore.Get(kv.prefix, kv.key)
	a.Nil(errGet)
	a.Equal(v, []byte(kv.value))
}

func Test2Close(t *testing.T) {
	inmemStore, err := NewBStore("", false)

	a := assert.New(t)
	a.Nil(err)

	errClose := inmemStore.Close()
	a.Nil(errClose)

	// test insert
	kv := KV{"prefix", "x", "y"}
	errSet := inmemStore.Set(kv)
	a.Error(errSet, "")
}

func Test3TTL(t *testing.T) {
	inmemStore, err := NewBStore("", false)
	defer inmemStore.Close()

	a := assert.New(t)
	a.Nil(err)

	kv := KV{"prefix", "x", "y"}
	ttl := 1

	errSet := inmemStore.SetTTL(kv, ttl)
	a.Nil(errSet)

	time.Sleep(time.Duration(ttl+1) * time.Second)
	_, errGet := inmemStore.Get(kv.prefix, kv.key)
	a.Error(errGet, "")
}
