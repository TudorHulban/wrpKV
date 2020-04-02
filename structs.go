package badgerwrap

import (
	badger "github.com/dgraph-io/badger/v2"
)

// KV is key value for the NoSQL DB.
type KV struct {
	prefix string
	key    string
	value  string
}

type bstore struct {
	*badger.DB
}
