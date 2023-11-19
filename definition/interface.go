package definition

import (
	"github.com/TudorHulban/kv"
	"github.com/TudorHulban/kv/kvbadger"
)

type KVStore interface {
	Set(kv.KV) error
	SetTTL(value kv.KV, secondsTTL uint) error
	SetAny([]byte, any) error
	SetAnyTTL(key []byte, value any, secondsTTL uint) error

	GetValueFor(key []byte) ([]byte, error)
	GetAnyByK([]byte, interface{}) error
	GetKVByPrefix(keyPrefix []byte) ([]*kv.KV, error)

	DeleteKVBy([]byte) error

	Close() error
}

var _ KVStore = &kvbadger.KVStore{}

// var _ KVStore = kvnuts.KVStore{}
