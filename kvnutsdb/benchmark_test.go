package kvnuts

import (
	"testing"
)

// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// Benchmark_InMemory-16    	  376618	      3747 ns/op	    1226 B/op	      23 allocs/op
func Benchmark_InMemory(b *testing.B) {
	b.ResetTimer()

	key := []byte("x")
	value := []byte("y")
	bucket := "A"

	store, _ := NewStoreInMemory(_segmentSizeTests)

	for i := 0; i < b.N; i++ {
		store.Set(bucket, key, value)
		store.GetValueByKey(bucket, key)
		store.DeleteKVByKey(bucket, key)
	}
}

func Benchmark_KeyInMemory(b *testing.B) {
	b.ResetTimer()

	key := []byte("x")
	value := []byte("y")
	bucket := "A"

	store, _ := NewStoreInMemory(_segmentSizeTests)

	for i := 0; i < b.N; i++ {
		store.Set(bucket, key, value)
		store.GetValueByKey(bucket, key)
		store.DeleteKVByKey(bucket, key)
	}
}
