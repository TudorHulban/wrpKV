package badgerwrap

import (
	badger "github.com/dgraph-io/badger/v2"
)

type db badger.DB

// KV is key value for the NoSQL DB.
type KV struct {
	key   []byte
	value []byte
}

// NewDB returns a database ready for use.
func NewDB(pDBFilePath string, pSyncRights bool) (*badger.DB, error) {
	var options badger.Options

	if len(pDBFilePath) == 0 {
		options = badger.DefaultOptions("").WithInMemory(true)
	} else {
		options = badger.DefaultOptions(pDBFilePath)
		options.WithSyncWrites(pSyncRights)
	}
	return badger.Open(options)
}

func (db *db) SetKV(pKV KV) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set(pKV.key, pKV.value)
	})
}
