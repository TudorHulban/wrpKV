package badgerwrap

import (
	"sync"
	"testing"

	"github.com/TudorHulban/loginfo"
	"github.com/stretchr/testify/assert"
)

func Test1ByPrefix(t *testing.T) {
	a := assert.New(t)

	l, errLog := loginfo.New(2)
	a.Nil(errLog)

	inmemStore, err := NewBStore("", false, l)
	defer inmemStore.Close()

	a.Nil(err)

	kPrefix := "prefix-"
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		kv1 := KV{kPrefix + "x1", "y1"}
		errSet := inmemStore.Set(kv1)
		a.Nil(errSet)

		wg.Done()
	}()

	go func() {
		kv2 := KV{kPrefix + "x2", "y2"}
		errSet := inmemStore.Set(kv2)
		a.Nil(errSet)

		wg.Done()
	}()

	go func() {
		kv3 := KV{kPrefix + "x3", "y3"}
		errSet := inmemStore.Set(kv3)
		a.Nil(errSet)

		wg.Done()
	}()

	wg.Wait()

	v, errGet := inmemStore.GetPrefix(kPrefix)
	a.Nil(errGet)
	a.Equal(len(v), 3)

	for i, v := range v {
		l.Info(i, v)
	}
}
