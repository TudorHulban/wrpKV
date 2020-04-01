package badgerwrap

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	inmemStore, err := NewBStore("", false)

	a := assert.New(t)
	a.Nil(err)

	kv := KV{[]byte("x"), []byte("y")}
	errSet := inmemStore.SetKV(kv)
	a.Nil(errSet)

	v, errGet := inmemStore.GetV(kv.key)
	a.Nil(errGet)
	a.Equal(v, kv.value)
}
