package stream

import "strings"

type Stream struct{}

func (alg Stream) GetEncryptionKey() string {
	return "3x3"
}
func (alg Stream) processText(txt string) string {
	var str strings.Builder
	for i := 1; i < len(txt); i += 2 {
		str.WriteByte(txt[i])
		str.WriteByte(txt[i-1])
	}
	return str.String()
}

func (alg Stream) Encrypt(msg string) string {
	return alg.processText(msg)
}

func (alg Stream) Decrypt(cip string) string {
	return alg.processText(cip)
}
