package badgerwrap

import (
	"testing"

	"github.com/TudorHulban/loginfo"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	l, errLog := loginfo.New(2)
	assert.Nil(t, errLog)

	inmemStore, err := NewBStoreInMem(l)
	assert.Nil(t, err)
	defer func() {
		assert.Nil(t, inmemStore.Close())
	}()

	kv := KV{[]byte("x"), []byte("y")}

	// test insert
	assert.Nil(t, inmemStore.Set(kv))

	v0, errGet := inmemStore.GetVByK(kv.key)
	assert.Nil(t, errGet)
	assert.Equal(t, v0, []byte(kv.value))

	// now delete the KV
	assert.Nil(t, inmemStore.DeleteKVByK(kv.key))
	v1, errGet := inmemStore.GetVByK(kv.key)
	assert.Error(t, errGet)
	assert.Nil(t, v1)
}
