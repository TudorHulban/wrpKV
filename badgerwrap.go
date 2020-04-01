package badgerwrap

import (
	badger "github.com/dgraph-io/badger/v2"
)

type bstore struct {
	*badger.DB
}

// NewBStore returns a type as per defined store interface.
func NewBStore(pDBFilePath string, pSyncRights bool) (store, error) {
	var options badger.Options

	if len(pDBFilePath) == 0 {
		options = badger.DefaultOptions("").WithInMemory(true)
	} else {
		options = badger.DefaultOptions(pDBFilePath)
		options.WithSyncWrites(pSyncRights)
	}
	result, errOpen := badger.Open(options)
	return bstore{
		result,
	}, errOpen
}

// SetKV sets key in store.
func (s bstore) SetKV(pKV KV) error {
	return s.Update(func(txn *badger.Txn) error {
		return txn.Set(pKV.key, pKV.value)
	})
}

// GetV fetches key from store.
func (s bstore) GetV(pKey []byte) ([]byte, error) {
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
