package kvbadger

import (
	"time"

	"github.com/TudorHulban/kv"
	"github.com/TudorHulban/kv/helpers"
	badger "github.com/dgraph-io/badger/v4"
)

func (s *KVStore) Set(_ string, item kv.KV) error {
	return s.Store.
		Update(
			func(txn *badger.Txn) error {
				return txn.Set(
					item.Key,
					item.Value,
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
		"",
		kv.KV{
			Key:   key,
			Value: encodedValue,
		},
	)
}

// SetTTL can be used for inserts and updates.
func (s *KVStore) SetTTL(bucket string, value kv.KV, secondsTTL uint) error {
	return s.Store.
		Update(
			func(txn *badger.Txn) error {
				return txn.
					SetEntry(
						badger.NewEntry(value.Key, value.Value).
							WithTTL(time.Second * time.Duration(secondsTTL)),
					)
			},
		)
}

// SetAnyTTL sets or updates key in store.
func (s *KVStore) SetAnyTTL(bucket string, key []byte, value any, secondsTTL uint) error {
	encodedValue, errEncode := helpers.Encode(value)
	if errEncode != nil {
		return errEncode
	}

	return s.SetTTL(
		"",
		kv.KV{
			Key:   key,
			Value: encodedValue,
		},

		secondsTTL,
	)
}
