package badgerwrap

// should be exported
type store interface {
	Set(pKV KV) error
	SetTTL(pKV KV, pTTLSecs int) error
	Get(string, string) ([]byte, error)
	Close() error
}
