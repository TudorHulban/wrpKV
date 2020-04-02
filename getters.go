package badgerwrap

import (
	badger "github.com/dgraph-io/badger/v2"
)

// GetV fetches key from store.
func (s bstore) Get(pKey []byte) ([]byte, error) {
	var result []byte

	errView := s.View(func(txn *badger.Txn) error {
		item, errGet := txn.Get(pKey)
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
