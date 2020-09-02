package badgerwrap

/*
File contains constructors for model.
Different constructors for disk sync options were written for better usability.
Passing logger as pointer as we might not want more than one logger.
*/

import (
	"github.com/TudorHulban/log"
	badger "github.com/dgraph-io/badger/v2"
	"github.com/pkg/errors"
)

// BStore Concentrates information defining a KV store.
type BStore struct {
	theLogger *log.LogInfo // logger needed only for package logging
	TheStore  *badger.DB
}

// NewBStoreDiskWSyncWrites returns a type containing a store that satisfies store interface.
func NewBStoreDiskWSyncWrites(dbFilePath string, extLogger *log.LogInfo) (*BStore, error) {
	dbBadger, errOpen := badger.Open(badger.DefaultOptions(dbFilePath))
	if errOpen != nil {
		return nil, errors.WithMessage(errOpen, "could not open passed file path in constructor")
	}

	return &BStore{
		theLogger: extLogger,
		TheStore:  dbBadger,
	}, errOpen
}

// NewBStoreDisk returns a type containing a store that satisfies store interface.
// No sync writes.
func NewBStoreDisk(dbFilePath string, extLogger *log.LogInfo) (*BStore, error) {
	dbBadger, errOpen := badger.Open(badger.DefaultOptions(dbFilePath).WithSyncWrites(false))
	if errOpen != nil {
		return nil, errors.WithMessage(errOpen, "could not open passed file path in constructor")
	}

	return &BStore{
		theLogger: extLogger,
		TheStore:  dbBadger,
	}, errOpen
}

// NewBStoreInMem Creates in memory Badger DB.
func NewBStoreInMem(extLogger *log.LogInfo) (*BStore, error) {
	dbBadger, errOpen := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if errOpen != nil {
		return nil, errors.WithMessage(errOpen, "error when creating in memory store")
	}

	return &BStore{
		theLogger: extLogger,
		TheStore:  dbBadger,
	}, nil
}

// NewBStoreInMemNoLogging Creates in memory Badger DB.
func NewBStoreInMemNoLogging() (*BStore, error) {
	options := badger.DefaultOptions("").WithInMemory(true).WithLogger(nil)

	dbBadger, errOpen := badger.Open(options)
	if errOpen != nil {
		return nil, errors.WithMessage(errOpen, "error when creating in memory store")
	}

	return &BStore{
		TheStore: dbBadger,
	}, nil
}

// Close closes the store.
func (s BStore) Close() error {
	return s.TheStore.Close()
}
