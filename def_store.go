package badgerwrap

type store interface {
	SetKV(pKV KV) error
	GetV(pKey []byte) ([]byte, error)
}
