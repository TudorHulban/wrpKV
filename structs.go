/*
	File contains structs belonging to badger wrap.
*/
package badgerwrap

import (
	"github.com/TudorHulban/loginfo"
	badger "github.com/dgraph-io/badger/v2"
)

// KV is key value for the NoSQL DB.
type KV struct {
	key   string
	value string
}

// With injected logger.
type bstore struct {
	theLogger loginfo.LogInfo
	b         *badger.DB
}
