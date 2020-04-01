package badgerwrap

// KV is key value for the NoSQL DB.
type KV struct {
	key   []byte
	value []byte
}

// should be exported
type store interface {
	SetKV(pKV KV) error
	GetV(pKey []byte) ([]byte, error)
}
