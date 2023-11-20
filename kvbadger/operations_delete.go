package kvbadger

import badger "github.com/dgraph-io/badger/v4"

func (s *KVStore) DeleteKVBy(_ string, key []byte) error {
	return s.Store.
		Update(
			func(txn *badger.Txn) error {
				return txn.Delete([]byte(key))
			})
}
