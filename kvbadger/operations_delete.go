package kvbadger

import badger "github.com/dgraph-io/badger/v4"

// DeleteKVBy Deletes KV by key.
func (s *KVStore) DeleteKVBy(key []byte) error {
	return s.Store.
		Update(
			func(txn *badger.Txn) error {
				return txn.Delete([]byte(key))
			})
}
