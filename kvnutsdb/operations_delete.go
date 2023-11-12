package kvnuts

import "github.com/nutsdb/nutsdb"

// DeleteKVByKey Deletes KV by key.
func (s *KVStore) DeleteKVByKey(bucket string, key []byte) error {
	return s.Store.
		Update(
			func(txn *nutsdb.Tx) error {
				return txn.Delete(bucket, []byte(key))
			},
		)
}
