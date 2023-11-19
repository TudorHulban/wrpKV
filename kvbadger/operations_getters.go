package kvbadger

import (
	"github.com/TudorHulban/kv"
	"github.com/TudorHulban/kv/helpers"
	badger "github.com/dgraph-io/badger/v4"
)

func (s *KVStore) GetValueFor(key []byte) ([]byte, error) {
	var res []byte

	errView := s.Store.
		View(func(txn *badger.Txn) error {
			value, errGet := txn.Get([]byte(key))
			if errGet != nil {
				return errGet
			}

			go func() {
				if s.logger != nil {
					s.logger.
						Debugf("size: %v, expires: %v",
							value.EstimatedSize(), value.ExpiresAt())
				}
			}()

			return value.Value(
				func(itemVals []byte) error {
					res = append(res, itemVals...)

					return nil
				})
		})

	return res, errView
}

func (s *KVStore) GetAnyByK(key []byte, result any) error {
	if helpers.CheckItemsArePointers(result) != -1 {
		return ErrNotAPointerType{}
	}

	encodedValue, errGet := s.GetValueFor(key)
	if errGet != nil {
		return errGet
	}

	return helpers.Decode([]byte(encodedValue), result)
}

func (s *KVStore) GetKVByPrefix(keyPrefix []byte) ([]*kv.KV, error) {
	var result []*kv.KV

	errView := s.Store.
		View(func(txn *badger.Txn) error {
			options := badger.DefaultIteratorOptions
			options.PrefetchSize = 10

			iterator := txn.NewIterator(options)
			defer iterator.Close()

			var errItem error

			for iterator.Seek(keyPrefix); iterator.ValidForPrefix(keyPrefix); iterator.Next() {
				item := iterator.Item()
				itemKey := item.Key()

				errItem = item.Value(
					func(itemValue []byte) error {
						go func() {
							if s.logger != nil {
								s.logger.Debugf("key=%s, value=%s", itemKey, itemValue)
							}
						}()

						result = append(result, &kv.KV{
							Key:   itemKey,
							Value: itemValue,
						})

						return nil
					})

				if errItem != nil {
					return errItem
				}
			}

			return nil
		})

	return result, errView
}
