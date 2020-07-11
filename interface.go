package badgerwrap

// Store is interface for interacting with KV stores.
type Store interface {
	Set(KV) error
	SetTTL(KV, int) error
	Get(string) ([]byte, error)
	GetPrefix(string) ([]KV, error)
	// Close closes the opened KV store.
	Close() error
}
