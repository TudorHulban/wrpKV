package kv

import (
	"github.com/TudorHulban/kv/kvbadger"
	kvnuts "github.com/TudorHulban/kv/kvnutsdb"
)

// KV is key value for the NoSQL DB.
type KV struct {
	Key   []byte
	Value []byte
}

// KVStore is interface for interacting with KV stores.
type KVStore interface {
	// Inserts or updates KV in store.
	Set(KV) error
	// Inserts or updates KV in store. Time To Live in seconds.
	SetTTL(KV, uint) error
	// Inserts or updates KV in store. Value is to be serialized structure.
	SetAny([]byte, interface{}) error
	// Inserts or updates KV in store. Value is to be serialized structure.
	SetAnyTTL([]byte, interface{}, uint) error
	// Returns value for passed key if found. If not found it returns empty slice and an error not nil.
	GetVByK([]byte) ([]byte, error)
	// Fills up the passed pointer value for passed key if found. If not found it returns an error not nil.
	GetAnyByK([]byte, interface{}) error
	// Returns a slice of KV if prefix found.
	// If not found it returns empty slice.
	GetKVByPrefix([]byte) ([]KV, error)
	// Deletes KV bassed on key.
	DeleteKVByK([]byte) error
	// Close closes the opened KV store.
	Close() error
}

var _ KVStore = kvbadger.KVStore{}
var _ KVStore = kvnuts.KVStore{}
