package kvnuts

import (
	"github.com/TudorHulban/kv/helpers"
	"github.com/xujiajun/nutsdb"
)

// Set sets or updates a key in store.
func (s *KVStore) Set(bucket string, key, value []byte) error {
	return s.Store.
		Update(
			func(txn *nutsdb.Tx) error {
				return txn.Put(bucket, key, value, 0)
			},
		)
}

// SetAny sets or updates key in store.
func (s *KVStore) SetAny(bucket string, key []byte, value any) error {
	encodedValue, errEncode := helpers.Encoder(value)
	if errEncode != nil {
		return errEncode
	}

	return s.Set(bucket, key, encodedValue)
}
