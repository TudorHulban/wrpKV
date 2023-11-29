package kv

import "fmt"

// KV is key value for the NoSQL DBs.
type KV struct {
	Key   []byte
	Value []byte
}

func (kv KV) String() string {
	return fmt.Sprintf("key: %s - value: %s", kv.Key, kv.Value)
}
