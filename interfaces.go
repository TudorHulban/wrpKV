package badgerwrap

// Customlogger is type to be used when injecting external logger.
type CustomLogger interface {
	Infof(string, ...interface{})
	Info(...interface{})
	Debugf(string, ...interface{})
	Debug(...interface{})
}

// Store is interface for interacting with KV stores.
type Store interface {
	Set(KV) error
	SetTTL(KV, int) error
	Get(string) ([]byte, error)
	GetPrefix(string) ([]KV, error)
	Close() error
}
