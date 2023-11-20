package kvnuts

import (
	"github.com/TudorHulban/kv/helpers"
	"github.com/nutsdb/nutsdb"
)

// GetValueByKey fetches value from store based on passed key.
func (s *KVStore) GetValueFor(bucket string, key []byte) ([]byte, error) {
	var res []byte

	errView := s.Store.
		View(
			func(txn *nutsdb.Tx) error {
				value, errGet := txn.Get(bucket, []byte(key))
				if errGet != nil {
					return errGet
				}

				res = value.Value

				return nil
			},
		)

	return res, errView
}

func (s *KVStore) GetAnyByK(bucket string, key []byte, result any) error {
	if helpers.CheckItemsArePointers(result) != -1 {
		return ErrNotAPointerType{}
	}

	encodedValue, errGet := s.GetValueFor(bucket, key)
	if errGet != nil {
		return errGet
	}

	return helpers.Decode([]byte(encodedValue), result)
}
