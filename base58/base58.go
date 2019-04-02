package base58

import (
	"math/big"
)

//go:generate go run genalphabet.go

var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

// Decode decodes a base58 string to a byte slice.
func Decode(b string) []byte {
	var numZeroes int

	answer := big.NewInt(0)
	j := big.NewInt(1)
	scratch := new(big.Int)

	for i := len(b) - 1; i >= 0; i-- {
		temp := base58[b[i]]

		if temp == 255 {
			return []byte("")
		}

		scratch.SetInt64(int64(temp))
		scratch.Mul(j, scratch)

		answer.Add(answer, scratch)

		j.Mul(j, bigRadix)
	}

	tempVal := answer.Bytes()

	for numZeroes = 0; numZeroes < len(b); numZeroes++ {
		if b[numZeroes] != alphabetIDx0 {
			break
		}
	}

	value := make([]byte, (numZeroes + len(tempVal)))
	copy(value[numZeroes:], tempVal)

	return value
}

// Encode encodes a byte slice to a base58 string.
func Encode(b []byte) string {
	x := new(big.Int)
	x.SetBytes(b)

	answer := make([]byte, 0, len(b)*136/100)

	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)

		answer = append(answer, alphabet[mod.Int64()])
	}

	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabetIDx0)
	}

	for i := 0; i < (len(answer))/2; i++ {
		answer[i], answer[len(answer)-1-i] = answer[len(answer)-1-i], answer[i]
	}

	return string(answer)
}
