package base58

import (
	"errors"

	"github.com/checksum0/go-cryptoutils/chainhash"
)

// ErrChecksum ...
var ErrChecksum = errors.New("checksum error")

// ErrInvalidFormat ...
var ErrInvalidFormat = errors.New("invalid format: version and/or checksum bytes missing")

func checksum(input []byte) (chcksum [4]byte) {
	hash := chainhash.SHA256dToBytes(input)
	copy(chcksum[:], hash[:4])

	return
}

// CheckEncode ...
func CheckEncode(input []byte, version byte) string {
	b := make([]byte, 0, len(input)+1+4)

	b = append(b, version)
	b = append(b, input[:]...)

	chcksum := checksum(b)

	b = append(b, chcksum[:]...)

	return Encode(b)
}

// CheckDecode ...
func CheckDecode(input string) (result []byte, version byte, err error) {
	var chcksum [4]byte

	decode := Decode(input)

	if len(decode) < 5 {
		return nil, 0, ErrInvalidFormat
	}

	version = decode[0]

	copy(chcksum[:], decode[len(decode)-4:])

	if checksum(decode[:len(decode)-4]) != chcksum {
		return nil, 0, ErrChecksum
	}

	payload := decode[1 : len(decode)-4]
	result = append(result, payload...)

	return
}
