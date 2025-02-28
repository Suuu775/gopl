package ex41

import (
	"crypto/sha256"

	"github.com/Suuu775/gopl/ch2/popcount"
)

func CountDiffBit(data1, data2 []byte) int {
	sum := 0
	xorarray := [32]byte{0}
	sm1 := sha256.Sum256(data1)
	sm2 := sha256.Sum256(data2)

	for i := 0; i < 32; i++ {
		xorarray[i] = sm1[i] ^ sm2[i]
		sum += popcount.PopCount(uint64(xorarray[i]))
	}
	return 256 - sum
}
