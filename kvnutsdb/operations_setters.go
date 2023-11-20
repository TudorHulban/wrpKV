package kvnuts

import (
	"github.com/TudorHulban/kv"
	"github.com/TudorHulban/kv/helpers"
	"github.com/nutsdb/nutsdb"
)

// Set sets or updates a key in store.
func (s *KVStore) Set(bucket string, item kv.KV) error {
	return s.Store.
		Update(
			func(txn *nutsdb.Tx) error {
				return txn.Put(
					bucket,
					item.Key,
					item.Value,
					0,
				)
			},
		)
}

// SetAny sets or updates key in store.
func (s *KVStore) SetAny(bucket string, key []byte, value any) error {
	encodedValue, errEncode := helpers.Encode(value)
	if errEncode != nil {
		return errEncode
	}

	return s.Set(
		bucket,
		kv.KV{
			Key:   key,
			Value: encodedValue,
		},
	)
}
