package badgerwrap

/*
	File contains structs belonging to badger wrap.
*/

import (
	"github.com/TudorHulban/loginfo"
	badger "github.com/dgraph-io/badger/v2"
)

// KV is key value for the NoSQL DB.
type KV struct {
	key   string
	value string
}

// BStore Concentrates information defining a KV store.
type BStore struct {
	TheLogger loginfo.LogInfo
	TheStore  *badger.DB
}
