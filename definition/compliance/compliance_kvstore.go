package compliance

import (
	"github.com/TudorHulban/kv/definition"
	"github.com/TudorHulban/kv/kvbadger"
	kvnuts "github.com/TudorHulban/kv/kvnutsdb"
)

var _ definition.KVStore = &kvbadger.KVStore{}

var _ definition.KVStore = &kvnuts.KVStore{}
