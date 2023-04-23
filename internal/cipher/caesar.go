package cipher

import "strings"

type Caesar struct{}

func (alg Caesar) GetEncryptionKey() string {
	return "2x2"
}

func (alg Caesar) Encrypt(msg string) string {
	var str strings.Builder
	for _, ch := range msg {
		str.WriteRune(ch + 5)
	}
	return str.String()
}

func (alg Caesar) Decrypt(cip string) string {
	var str strings.Builder
	for _, ch := range cip {
		str.WriteRune(ch - 5)
	}
	return str.String()
}
