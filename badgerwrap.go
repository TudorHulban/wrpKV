package badgerwrap

import (
	badger "github.com/dgraph-io/badger/v2"
)

type store struct {
	*badger.DB
}

// KV is key value for the NoSQL DB.
type KV struct {
	key   []byte
	value []byte
}

// NewStore returns a database ready for use.
func NewStore(pDBFilePath string, pSyncRights bool) (*store, error) {
	var options badger.Options

	if len(pDBFilePath) == 0 {
		options = badger.DefaultOptions("").WithInMemory(true)
	} else {
		options = badger.DefaultOptions(pDBFilePath)
		options.WithSyncWrites(pSyncRights)
	}
	result, errOpen := badger.Open(options)
	return &store{
		result,
	}, errOpen
}

func (s store) SetKV(pKV KV) error {
	return s.Update(func(txn *badger.Txn) error {
		return txn.Set(pKV.key, pKV.value)
	})
}
