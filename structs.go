package badgerwrap

import (
	badger "github.com/dgraph-io/badger/v2"
)

// KV is key value for the NoSQL DB.
type KV struct {
	key   []byte
	value []byte
}

type bstore struct {
	*badger.DB
}
