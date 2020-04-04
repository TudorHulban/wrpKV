package badgerwrap

import (
	"time"

	badger "github.com/dgraph-io/badger/v2"
)

// SetKV sets or updates key in store.
func (s bstore) Set(pKV KV) error {
	return s.b.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(pKV.key), []byte(pKV.value))
	})
}

// setKVTTL can be used for inserts and updates. Time To Live in seconds.
func (s bstore) SetTTL(pKV KV, pTTLSecs int) error {
	return s.b.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(pKV.key), []byte(pKV.value)).WithTTL(time.Second * time.Duration(pTTLSecs))
		errSet := txn.SetEntry(entry)
		return errSet
	})
}
