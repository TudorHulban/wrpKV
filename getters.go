package badgerwrap

import (
	badger "github.com/dgraph-io/badger/v2"
)

// GetV fetches key from store.
func (s bstore) Get(pPrefix, pKey string) ([]byte, error) {
	var result []byte

	errView := s.View(func(txn *badger.Txn) error {
		item, errGet := txn.Get([]byte(pPrefix + pKey))
		if errGet != nil {
			return errGet
		}
		//log.Println("size:", item.EstimatedSize(), item.ExpiresAt())
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
	prefix := []byte(pKeyPrefix)

	errView := s.View(func(txn *badger.Txn) error {
		iterator := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iterator.Close()

		var errItem error
		for {
			iterator.Seek(prefix)
			item := iterator.Item()
			//log.Println("item:", item)
			//k := item.Key()

			//log.Println("key:", string(k))

			errItem := item.Value(func(itemVals []byte) error {
				//log.Printf("key=%s, value=%s\n", k, itemVals)
				return nil
			})
			return errItem

			if !iterator.ValidForPrefix(prefix) {
				break
			}
			iterator.Next()
		}
		return errItem
	})
	return result, errView
}
