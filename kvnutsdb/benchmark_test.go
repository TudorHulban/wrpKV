package kvnuts

import (
	"testing"

	"github.com/TudorHulban/kv"
)

// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// Benchmark_InMemory-16    	  376618	      3747 ns/op	    1226 B/op	      23 allocs/op
func Benchmark_InMemory(b *testing.B) {
	key := []byte("x")
	value := []byte("y")
	bucket := "A"

	store, errCr := NewStoreInMemory(_segmentSizeTests)
	b.Log(errCr)

	if errCr != nil {
		b.FailNow()
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		store.Set(
			bucket,
			kv.KV{
				Key:   key,
				Value: value,
			},
		)
		store.GetValueFor(bucket, key)
		store.DeleteKVBy(bucket, key)
	}
}

func Benchmark_KeyInMemory(b *testing.B) {
	key := []byte("x")
	value := []byte("y")
	bucket := "A"

	store, errCr := NewStoreInMemory(_segmentSizeTests)
	b.Log(errCr)

	if errCr != nil {
		b.FailNow()
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		store.Set(bucket, kv.KV{
			Key:   key,
			Value: value,
		})
		store.GetValueFor(bucket, key)
		store.DeleteKVBy(bucket, key)
	}
}
