package badgerwrap

import (
	"testing"
	"time"

	"github.com/TudorHulban/loginfo"
	"github.com/stretchr/testify/assert"
)

func Test1Set(t *testing.T) {
	a := assert.New(t)

	l, errLog := loginfo.New(2)
	a.Nil(errLog)

	inmemStore, err := NewBStoreInMem(l)
	defer inmemStore.Close()

	a.Nil(err)

	kPrefix := "prefix-"
	kv := KV{kPrefix + "x", "y"}

	// test insert
	errSet := inmemStore.Set(kv)
	a.Nil(errSet)

	// test update
	kv.value = "z"
	errUpdate := inmemStore.Set(kv)
	a.Nil(errUpdate)

	v, errGet := inmemStore.GetKVByK(kv.key)
	a.Nil(errGet)
	a.Equal(v, []byte(kv.value))
}

func Test2Close(t *testing.T) {
	a := assert.New(t)

	l, errLog := loginfo.New(2)
	a.Nil(errLog)

	inmemStore, err := NewBStoreInMem(l)

	a.Nil(err)

	errClose := inmemStore.Close()
	a.Nil(errClose)

	// test insert
	kv := KV{"prefix-x", "y"}
	errSet := inmemStore.Set(kv)
	a.Error(errSet, "")
}

func Test3TTL(t *testing.T) {
	a := assert.New(t)

	l, errLog := loginfo.New(2)
	a.Nil(errLog)

	inmemStore, err := NewBStoreInMem(l)
	defer inmemStore.Close()

	a.Nil(err)

	kPrefix := "prefix-"
	kv := KV{kPrefix + "x", "y"}
	ttl := 1

	errSet := inmemStore.SetTTL(kv, uint8(ttl))
	a.Nil(errSet)

	time.Sleep(time.Duration(ttl+1) * time.Second)
	_, errGet := inmemStore.GetKVByK(kv.key)
	a.Error(errGet, "")
}
