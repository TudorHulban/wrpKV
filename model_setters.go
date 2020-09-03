package badgerwrap

import (
	"time"

	badger "github.com/dgraph-io/badger/v2"
)

// Set sets or updates key in store.
func (s BStore) Set(theKV KV) error {
	return s.TheStore.Update(func(txn *badger.Txn) error {
		return txn.Set(theKV.Key, theKV.Value)
	})
}

// SetAny sets or updates key in store.
func (s BStore) SetAny(theKey []byte, theValue interface{}) error {
	v, errEncode := anyEncoder(theValue)
	if errEncode != nil {
		return errEncode
	}

	return s.Set(KV{
		Key:   theKey,
		Value: v,
	})
}

// SetAnyTTL sets or updates key in store.
func (s BStore) SetAnyTTL(theKey []byte, theValue interface{}, ttlSecs uint) error {
	v, errEncode := anyEncoder(theValue)
	if errEncode != nil {
		return errEncode
	}

	return s.SetTTL(KV{
		Key:   theKey,
		Value: v,
	}, ttlSecs)
}

// SetTTL can be used for inserts and updates. Time To Live in seconds.
func (s BStore) SetTTL(theKV KV, ttlSecs uint) error {
	return s.TheStore.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry(theKV.Key, theKV.Value).WithTTL(time.Second * time.Duration(ttlSecs))
		return txn.SetEntry(entry)
	})
}
