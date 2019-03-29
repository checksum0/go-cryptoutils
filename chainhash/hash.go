package chainhash

import (
	"encoding/hex"
	"fmt"
)

// HashSize ...
const HashSize = 32

// MaxHashStringSize ...
const MaxHashStringSize = HashSize * 2

// ErrHashStrSize ...
var ErrHashStrSize = fmt.Errorf("max hash string length is %v bytes", MaxHashStringSize)

// Hash ...
type Hash [HashSize]byte

// String ...
func (hash Hash) String() string {
	for i := 0; i < HashSize/2; i++ {
		hash[i], hash[HashSize-1-i] = hash[HashSize-1-i], hash[i]
	}
	return hex.EncodeToString(hash[:])
}

// CopyBytes ...
func (hash *Hash) CopyBytes() []byte {
	newHash := make([]byte, HashSize)
	copy(newHash, hash[:])

	return newHash
}

// SetBytes ...
func (hash *Hash) SetBytes(newHash []byte) error {
	hlen := len(newHash)
	if hlen != HashSize {
		return fmt.Errorf("invalid hash length of %v, %v needed", hlen, HashSize)
	}
	copy(hash[:], newHash)

	return nil
}

// Compare ...
func (hash *Hash) Compare(target *Hash) int {
	if hash == nil && target == nil {
		return 0
	} else if hash != nil && target == nil {
		return 1
	} else if hash == nil && target != nil {
		return -1
	}

	h := *hash
	t := *target
	for i := 0; i < len(h); i++ {
		a := h[len(h)-1-i]
		b := t[len(t)-1-i]
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
	}

	return 0
}

// IsEqual ...
func (hash *Hash) IsEqual(target *Hash) bool {
	if hash == nil && target == nil {
		return true
	} else if hash == nil || target == nil {
		return false
	}

	return *hash == *target
}

// NewHashFromBytes ...
func NewHashFromBytes(newHash []byte) (*Hash, error) {
	var ret Hash
	err := ret.SetBytes(newHash)
	if err != nil {
		return nil, err
	}

	return &ret, err
}

// NewHashFromString ...
func NewHashFromString(newHash string) (*Hash, error) {
	ret := new(Hash)
	err := Decode(ret, newHash)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// Decode ...
func Decode(dst *Hash, src string) error {
	if len(src) > MaxHashStringSize {
		return ErrHashStrSize
	}

	var srcBytes []byte
	if len(src)%2 == 0 {
		srcBytes = []byte(src)
	} else {
		srcBytes = make([]byte, 1+len(src))
		srcBytes[0] = '0'
		copy(srcBytes[1:], src)
	}

	var reversed Hash
	_, err := hex.Decode(reversed[HashSize-hex.DecodedLen(len(srcBytes)):], srcBytes)
	if err != nil {
		return err
	}

	for i, b := range reversed[:HashSize/2] {
		dst[i], dst[HashSize-1-i] = reversed[HashSize-1-i], b
	}

	return nil
}
