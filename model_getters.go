package badgerwrap

import (
	badger "github.com/dgraph-io/badger/v2"
)

// GetV fetches key from store.
func (s bstore) Get(pKey string) ([]byte, error) {
	var result []byte

	errView := s.b.View(func(txn *badger.Txn) error {
		item, errGet := txn.Get([]byte(pKey))
		if errGet != nil {
			return errGet
		}
		s.theLogger.Debugf("size: %s, expires: %s", item.EstimatedSize(), item.ExpiresAt())

		errItem := item.Value(func(itemVals []byte) error {
			result = append(result, itemVals...)
			return nil
		})
		return errItem
	})
	return result, errView
}

// GetPrefix in case it does not find keys, returns first key in store.
func (s bstore) GetPrefix(pKeyPrefix string) ([]KV, error) {
	var result []KV

	errView := s.b.View(func(txn *badger.Txn) error {
		options := badger.DefaultIteratorOptions
		options.PrefetchSize = 10

		iterator := txn.NewIterator(options)
		defer iterator.Close()

		prefix := []byte(pKeyPrefix)
		var errItem error

		for iterator.Seek(prefix); iterator.ValidForPrefix(prefix); iterator.Next() {
			item := iterator.Item()
			k := item.Key()

			errItem = item.Value(func(itemValue []byte) error {
				s.theLogger.Debugf("key=%s, value=%s\n", k, itemValue)

				kv := KV{string(k), string(itemValue)}
				result = append(result, kv)
				return nil
			})
			if errItem != nil {
				return errItem
			}
		}
		return errItem
	})
	return result, errView
}
