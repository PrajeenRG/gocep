package cipher

import "strings"

type Block struct {
	Key string
}

func (alg Block) GetEncryptionKey() string {
	return "1x1"
}

func (alg Block) getKey(pos int) byte {
	return alg.Key[pos%len(alg.Key)]
}

func (alg Block) Encrypt(msg string) string {
	var str strings.Builder
	for i := 0; i < len(msg); i++ {
		str.WriteByte((msg[i] ^ alg.getKey(i)) - alg.getKey(i))
	}
	return str.String()
}

func (alg Block) Decrypt(cip string) string {
	var str strings.Builder
	for i := 0; i < len(cip); i++ {
		str.WriteByte((cip[i] + alg.getKey(i)) ^ alg.getKey(i))
	}
	return str.String()
}
