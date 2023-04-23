package cipher

import (
	"strconv"
	"strings"
)

type Transpose struct {
	Dimen int
}

func (alg Transpose) GetEncryptionKey() string {
	return "4x4"
}

func (alg Transpose) getBlocks(txt string) [][][]byte {
	if bal := len(txt) % (alg.Dimen * alg.Dimen); bal != 0 {
		txt += strings.Repeat(strconv.Itoa(0), bal)
	}
	var blocks [][][]byte
	for i := 0; i < len(txt)/(alg.Dimen*alg.Dimen); i++ {
		block := make([][]byte, alg.Dimen)
		for j := 0; j < alg.Dimen; j++ {
			start := i + (j * alg.Dimen)
			end := i + ((j + 1) * alg.Dimen)
			block[j] = []byte(txt[start:end])
		}
		blocks = append(blocks, block)
	}
	return blocks
}

func (alg Transpose) matTranspose(mat [][]byte) [][]byte {
	for i := 0; i < alg.Dimen; i++ {
		for j := 0; j < i; j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}
	return mat
}

func (alg Transpose) Encrypt(msg string) string {
	var str strings.Builder
	blocks := alg.getBlocks(msg)
	for _, block := range blocks {
		block := alg.matTranspose(block)
		for _, row := range block {
			str.Write(row)
		}
	}
	return str.String()
}

func (alg Transpose) Decrypt(cip string) string {
	var str strings.Builder
	blocks := alg.getBlocks(cip)
	for _, block := range blocks {
		block := alg.matTranspose(block)
		for _, row := range block {
			str.Write(row)
		}
	}
	return str.String()
}
