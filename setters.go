/*
	Provides essential operations for interacting with Badger.

	To add more info.
*/
package badgerwrap

import (
	"time"

	badger "github.com/dgraph-io/badger/v2"
)

// SetKV sets or updates key in store.
func (s bstore) Set(pKV KV) error {
	return s.Update(func(txn *badger.Txn) error {
		return txn.Set(pKV.key, pKV.value)
	})
}

// setKVTTL can be used for inserts and updates. Time To Live in seconds.
func (s bstore) SetTTL(pKV KV, pTTLSecs int) error {
	return s.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry(pKV.key, pKV.value).WithTTL(time.Second * time.Duration(pTTLSecs))
		errSet := txn.SetEntry(entry)
		return errSet
	})
}
