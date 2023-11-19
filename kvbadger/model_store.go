package kvbadger

import (
	"fmt"

	"github.com/TudorHulban/log"
	badger "github.com/dgraph-io/badger/v4"
)

type KVStore struct {
	logger *log.Logger
	Store  *badger.DB
}

// NewBStoreDiskWSyncWrites returns a type containing a store that satisfies store interface.
func NewBStoreDiskWSyncWrites(dbFilePath string, l *log.Logger) (*KVStore, error) {
	db, errOpen := badger.Open(
		badger.DefaultOptions(dbFilePath),
	)
	if errOpen != nil {
		return nil,
			fmt.Errorf("open passed file %s in NewBStoreDiskWSyncWrites:%w",
				dbFilePath, errOpen)
	}

	return &KVStore{
			logger: l,
			Store:  db,
		},
		nil
}

// NewBStoreDisk returns a type containing a store that satisfies store interface.
// No sync writes.
func NewBStoreDisk(dbFilePath string, l *log.Logger) (*KVStore, error) {
	db, errOpen := badger.Open(
		badger.DefaultOptions(dbFilePath).
			WithSyncWrites(false))
	if errOpen != nil {
		return nil,
			fmt.Errorf("open passed file %s in NewBStoreDisk:%w",
				dbFilePath, errOpen)
	}

	return &KVStore{
			logger: l,
			Store:  db,
		},
		nil
}

func NewBStoreInMemory(extLogger *log.Logger) (*KVStore, error) {
	db, errOpen := badger.Open(
		badger.DefaultOptions("").
			WithInMemory(true))
	if errOpen != nil {
		return nil,
			fmt.Errorf("when creating in memory store:%w", errOpen)
	}

	return &KVStore{
			logger: extLogger,
			Store:  db,
		},
		nil
}

// NewBStoreInMemoryNoLogging Creates in memory Badger DB.
// No protection for writing to nil logger field.
func NewBStoreInMemoryNoLogging() (*KVStore, error) {
	options := badger.DefaultOptions("").
		WithInMemory(true).
		WithLogger(nil)

	db, errOpen := badger.Open(options)
	if errOpen != nil {
		return nil,
			fmt.Errorf("when creating in memory store:%w", errOpen)
	}

	return &KVStore{
			Store: db,
		},
		nil
}

func (s *KVStore) Close() error {
	return s.Store.Close()
}
