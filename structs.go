/*
	File contains structs belonging to badger wrap.
*/
package badgerwrap

import (
	badger "github.com/dgraph-io/badger/v2"
)

// KV is key value for the NoSQL DB.
type KV struct {
	key   string
	value string
}

type bstore struct {
	theLogger CustomLogger
	b         *badger.DB
}
