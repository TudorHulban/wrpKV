package badger

import (
	"github.com/TudorHulban/badgerwrap/badger/kv"
	badger "github.com/dgraph-io/badger/v2"
)

// GetVByK fetches value from store based on passed key.
func (s BStore) GetVByK(theK []byte) ([]byte, error) {
	var result []byte

	errView := s.Store.View(func(txn *badger.Txn) error {
		item, errGet := txn.Get([]byte(theK))
		if errGet != nil {
			return errGet
		}
		s.logger.Debugf("size: %v, expires: %v", item.EstimatedSize(), item.ExpiresAt())

		return item.Value(func(itemVals []byte) error {
			result = append(result, itemVals...)
			return nil
		})
	})
	return result, errView
}

// GetAnyByK fetches value and injects it into passed pointer type.
func (s BStore) GetAnyByK(theK []byte, decodeInTo interface{}) error {
	v, errGet := s.GetVByK(theK)
	if errGet != nil {
		return errGet
	}
	return anyDecoder([]byte(v), decodeInTo)
}

// GetKVByPrefix in case it does not find keys, returns first key in store.
func (s BStore) GetKVByPrefix(theKPrefix []byte) ([]kv.KV, error) {
	var result []KV

	errView := s.Store.View(func(txn *badger.Txn) error {
		options := badger.DefaultIteratorOptions
		options.PrefetchSize = 10

		iterator := txn.NewIterator(options)
		defer iterator.Close()

		prefix := theKPrefix
		var errItem error

		for iterator.Seek(prefix); iterator.ValidForPrefix(prefix); iterator.Next() {
			item := iterator.Item()
			k := item.Key()

			errItem = item.Value(func(itemValue []byte) error {
				s.logger.Debugf("key=%s, value=%s\n", k, itemValue)

				result = append(result, KV{
					k,
					itemValue,
				})
				return nil
			})
			// eraly exit if any error
			if errItem != nil {
				return errItem
			}
		}
		return errItem
	})
	return result, errView
}
