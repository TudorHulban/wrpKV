package kvbadger

import (
	"time"

	"github.com/TudorHulban/kv/helpers"
	badger "github.com/dgraph-io/badger/v4"
)

func (s *KVStore) Set(key, value []byte) error {
	return s.Store.
		Update(
			func(txn *badger.Txn) error {
				return txn.Set(key, value)
			},
		)
}

// SetKV sets or updates a key value in store.
func (s *KVStore) SetKV(kv KV) error {
	return s.Set(kv.Key, kv.Value)
}

// SetAny sets or updates key in store.
func (s *KVStore) SetAny(key []byte, value any) error {
	encodedValue, errEncode := helpers.Encode(value)
	if errEncode != nil {
		return errEncode
	}

	return s.Set(key, encodedValue)
}

// SetTTL can be used for inserts and updates.
func (s *KVStore) SetTTL(key, value []byte, secondsTTL uint) error {
	return s.Store.
		Update(
			func(txn *badger.Txn) error {
				return txn.
					SetEntry(
						badger.NewEntry(key, value).
							WithTTL(time.Second * time.Duration(secondsTTL)),
					)
			})
}

// SetAnyTTL sets or updates key in store.
func (s *KVStore) SetAnyTTL(key []byte, value any, secondsTTL uint) error {
	encodedValue, errEncode := helpers.Encode(value)
	if errEncode != nil {
		return errEncode
	}

	return s.SetTTL(key, encodedValue, secondsTTL)
}
