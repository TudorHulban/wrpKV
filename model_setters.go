package badgerwrap

import (
	"time"

	badger "github.com/dgraph-io/badger/v2"
)

// SetKV sets or updates key in store.
func (s bstore) Set(theKV KV) error {
	return s.b.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(theKV.key), []byte(theKV.value))
	})
}

// setKVTTL can be used for inserts and updates. Time To Live in seconds.
func (s bstore) SetTTL(theKV KV, ttlSecs int) error {
	return s.b.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(theKV.key), []byte(theKV.value)).WithTTL(time.Second * time.Duration(ttlSecs))
		errSet := txn.SetEntry(entry)
		return errSet
	})
}
