package cipher

import "strings"

type Xor struct {
	Key int
}

func (alg Xor) GetEncryptionKey() string {
	return "6x6"
}

func (alg Xor) processText(txt string) string {
	var str strings.Builder
	for _, ch := range txt {
		str.WriteRune(rune(alg.Key) ^ ch)
	}
	return str.String()
}

func (alg Xor) Encrypt(msg string) string {
	return alg.processText(msg) + alg.GetEncryptionKey()
}

func (alg Xor) Decrypt(cip string) string {
	return alg.processText(cip)
}
