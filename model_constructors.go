package badgerwrap

/*
File contains constructors for model.
Different constructors for disk sync options were written for better usability.
*/

import (
	"github.com/TudorHulban/loginfo"
	badger "github.com/dgraph-io/badger/v2"
)

// NewBStoreDiskWSyncWrites returns a type containing a store that satisfies store interface.
func NewBStoreDiskWSyncWrites(dbFilePath string, extLogger loginfo.LogInfo) (BStore, error) {
	dbBadger, errOpen := badger.Open(badger.DefaultOptions(dbFilePath))
	if errOpen != nil {
		return BStore{
			theLogger: extLogger,
			TheStore:  nil,
		}, errOpen
	}

	return BStore{
		theLogger: extLogger,
		TheStore:  dbBadger,
	}, errOpen
}

// NewBStoreDisk returns a type containing a store that satisfies store interface.
// No sync writes.
func NewBStoreDisk(dbFilePath string, extLogger loginfo.LogInfo) (BStore, error) {
	options := badger.DefaultOptions(dbFilePath)
	options.WithSyncWrites(false)

	dbBadger, errOpen := badger.Open(options)
	if errOpen != nil {
		return BStore{
			theLogger: extLogger,
			TheStore:  nil,
		}, errOpen
	}

	return BStore{
		theLogger: extLogger,
		TheStore:  dbBadger,
	}, errOpen
}

// NewBStoreInMem Creates in memory Badger DB.
func NewBStoreInMem(extLogger loginfo.LogInfo) (BStore, error) {
	dbBadger, errOpen := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if errOpen != nil {
		return BStore{
			theLogger: extLogger,
			TheStore:  nil,
		}, errOpen
	}

	return BStore{
		theLogger: extLogger,
		TheStore:  dbBadger,
	}, nil
}

// Close closes the store.
func (s BStore) Close() error {
	return s.TheStore.Close()
}
