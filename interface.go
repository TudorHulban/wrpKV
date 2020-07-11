package badgerwrap

/*
File contains interface for interacting with KV stores and KV struct.
*/

// KV is key value for the NoSQL DB.
type KV struct {
	key   []byte
	value []byte
}

// Store is interface for interacting with KV stores.
type Store interface {
	// Inserts or updates KV in store.
	Set(KV) error
	// Inserts or updates KV in store. Time To Live in seconds.
	SetTTL(KV, int) error
	// Returns value of key if found. If not found it returns empty slice
	// and an error not nil.
	GetVByK([]byte) ([]byte, error)
	// Returns a slice of KV if prefix found.
	// If not found it returns the first key in store.
	GetKVByPrefix([]byte) ([]KV, error)
	// Close closes the opened KV store.
	Close() error
}
