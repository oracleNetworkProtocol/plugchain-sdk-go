package store

type AES struct{}

func (AES) Encrypt(text string, key string) (string, error) {
	return "", nil
}

func (AES) Decrypt(cryptoText string, key string) (string, error) {
	return "", nil
}
