package badgerwrap

// logger is type to be used when injecting external logger.
type logger interface {
	Debugf(string, ...interface{})
}

// Store is interface for interacting with KV stores.
type Store interface {
	Set(KV) error
	SetTTL(KV, int) error
	Get(string) ([]byte, error)
	GetPrefix(string) ([]KV, error)
	Close() error
}
