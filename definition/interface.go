package definition

import (
	"github.com/TudorHulban/kv"
	"github.com/TudorHulban/kv/kvbadger"
	kvnuts "github.com/TudorHulban/kv/kvnutsdb"
)

type KVStore interface {
	// Inserts or updates KV in store.
	Set(kv.KV) error
	// Inserts or updates KV in store. Time To Live in seconds.
	SetTTL(kv.KV, uint) error
	// Inserts or updates KV in store. Value is to be serialized structure.
	SetAny([]byte, interface{}) error
	// Inserts or updates KV in store. Value is to be serialized structure.
	SetAnyTTL([]byte, interface{}, uint) error

	GetValueFor(key []byte) ([]byte, error)
	GetAnyByK([]byte, interface{}) error
	GetKVByPrefix(keyPrefix []byte) ([]kv.KV, error)

	DeleteKVBy([]byte) error

	Close() error
}

var _ KVStore = &kvbadger.KVStore{}
var _ KVStore = kvnuts.KVStore{}
