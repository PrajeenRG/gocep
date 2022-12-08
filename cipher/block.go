package block

import "strings"

type Block struct {
	key string
}

func (alg Block) getKey(pos int) byte {
	return alg.key[pos%len(alg.key)]
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
