package cardcrypter

type CardNumber [16]byte

type Card struct {
	ID     string
	Number CardNumber
}

type Crypter interface {
	Encrypt(cards []Card, key []byte) ([]string, error)
}

type crypterImpl struct{}

func New(opts ...CrypterOption) *crypterImpl {
	panic("not implemented")
}

type CrypterOption func(*crypterImpl)

func WithWorkers(workers int) CrypterOption {
	panic("not implemented")
}

func (c *crypterImpl) Encrypt(cards []Card, key []byte) ([]string, error) {
	panic("not implemented")
}
