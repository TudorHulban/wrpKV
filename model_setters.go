package badgerwrap

import (
	"time"

	badger "github.com/dgraph-io/badger/v2"
)

// Set sets or updates key in store.
func (s BStore) Set(theKV KV) error {
	return s.TheStore.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(theKV.key), []byte(theKV.value))
	})
}

// SetTTL can be used for inserts and updates. Time To Live in seconds.
func (s BStore) SetTTL(theKV KV, ttlSecs uint8) error {
	return s.TheStore.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(theKV.key), []byte(theKV.value)).WithTTL(time.Second * time.Duration(ttlSecs))
		errSet := txn.SetEntry(entry)
		return errSet
	})
}
