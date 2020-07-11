package badgerwrap

import (
	"sync"
	"testing"

	"github.com/TudorHulban/loginfo"
	"github.com/stretchr/testify/assert"
)

func Test1ByPrefix(t *testing.T) {
	l, errLog := loginfo.New(2)
	assert.Nil(t, errLog)

	inmemStore, err := NewBStoreInMem(l)
	assert.Nil(t, err)
	defer func() {
		assert.Nil(t, inmemStore.Close())
	}()

	kPrefix := "prefix-"
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		kv1 := KV{[]byte(kPrefix + "x1"), []byte("y1")}
		errSet := inmemStore.Set(kv1)
		assert.Nil(t, errSet)

		wg.Done()
	}()

	go func() {
		kv2 := KV{[]byte(kPrefix + "x2"), []byte("y2")}
		errSet := inmemStore.Set(kv2)
		assert.Nil(t, errSet)

		wg.Done()
	}()

	go func() {
		kv3 := KV{[]byte(kPrefix + "x3"), []byte("y3")}
		errSet := inmemStore.Set(kv3)
		assert.Nil(t, errSet)

		wg.Done()
	}()

	wg.Wait()

	v, errGet := inmemStore.GetKVByPrefix([]byte(kPrefix))
	assert.Nil(t, errGet)
	assert.Equal(t, len(v), 3)

	for i, v := range v {
		l.Info(i, v)
	}
}
