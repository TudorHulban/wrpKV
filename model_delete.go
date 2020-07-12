package badgerwrap

import badger "github.com/dgraph-io/badger/v2"

// DeleteKVByK Delete KV by key.
func (s BStore) DeleteKVByK(theK []byte) error {
	return s.TheStore.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(theK))
	})
}
