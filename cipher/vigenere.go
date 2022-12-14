package cipher

import "strings"

type Vigenere struct{}

func (alg Vigenere) GetEncryptionKey() string {
	return "5x5"
}

func (alg Vigenere) Encrypt(msg string) string {
	var str strings.Builder
	var ch rune
	for _, val := range msg {
		switch {
		case val > 'z':
			ch = val
		case val >= 'A' && val <= 'Z':
			ch = val - 'A' + 1
		case val >= 'a' && val <= 'z':
			ch = 26 + val - 'a' + 1
		case val >= '0' && val <= '9':
			ch = 52 + val - '0' + 1
		case val > 'Z':
			ch = val + 26
		case val < '0':
			ch = val + 62
		case val > '9':
			ch = val + 52
		}
		str.WriteRune(ch)
	}
	return str.String()
}

func (alg Vigenere) Decrypt(cip string) string {
	var str strings.Builder
	var ch rune
	for _, val := range cip {
		switch {
		case val > 'z':
			ch = val
		case val >= 1 && val <= 26:
			ch = val + 'A' - 1
		case val >= 27 && val <= 52:
			ch = val + 'a' - 27
		case val >= 53 && val <= 62:
			ch = val + '0' - 53
		case val < 110:
			ch = val - 62
		case val < 117:
			ch = val - 52
		case val < 123:
			ch = val - 26
		}
		str.WriteRune(ch)
	}
	return str.String()
}
