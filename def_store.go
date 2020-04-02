package badgerwrap

// should be exported
type store interface {
	Set(pKV KV) error
	SetTTL(pKV KV, pTTLSecs int) error
	Get(pKey []byte) ([]byte, error)
	Close() error
}
