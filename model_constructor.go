package badgerwrap

import (
	"github.com/TudorHulban/loginfo"
	badger "github.com/dgraph-io/badger/v2"
)

// NewBStore returns a type as per defined store interface. This way only the contract is exposed.
func NewBStore(dbFilePath string, syncRights bool, extLogger loginfo.LogInfo) (Store, error) {
	var options badger.Options

	if len(dbFilePath) == 0 {
		options = badger.DefaultOptions("").WithInMemory(true)
	} else {
		options = badger.DefaultOptions(pDBFilePath)
		options.WithSyncWrites(syncRights)
	}
	result, errOpen := badger.Open(options)

	return bstore{
		theLogger: extLogger,
		b:         result,
	}, errOpen
}

// Close closes the store.
func (s bstore) Close() error {
	return s.b.Close()
}
