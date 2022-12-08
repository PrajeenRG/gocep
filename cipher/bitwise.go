package bitwise

import "strings"

type Bitwise struct{}

func (alg Bitwise) GetEncryptionKey() string {
	return "0x0"
}

func (alg Bitwise) processText(txt string) string {
	var str strings.Builder
	for _, ch := range txt {
		str.WriteRune(^ch)
	}
	return str.String()
}

func (alg Bitwise) Encrypt(msg string) string {
	return alg.processText(msg) + alg.GetEncryptionKey()
}

func (alg Bitwise) Decrypt(cip string) string {
	return alg.processText(cip)
}
