package definition

import (
	"github.com/TudorHulban/kv"
)

type KVStore interface {
	Set(bucket string, item kv.KV) error
	SetTTL(bucket string, value kv.KV, secondsTTL uint) error
	SetAny(bucket string, key []byte, value any) error
	SetAnyTTL(bucket string, key []byte, value any, secondsTTL uint) error

	GetValueFor(bucket string, key []byte) ([]byte, error)
	GetAnyByK(bucket string, key []byte, result any) error
	GetKVByPrefix(bucket string, keyPrefix []byte) ([]*kv.KV, error)

	DeleteKVBy(bucket string, key []byte) error

	Close() error
}
