package blockchain

import (
	"math"

	"github.com/checksum0/go-cryptoutils/chainhash"
)

// HashMerkleBranch ...
func HashMerkleBranch(left *chainhash.Hash, right *chainhash.Hash) *chainhash.Hash {
	var hashes [chainhash.HashSize * 2]byte
	copy(hashes[:chainhash.HashSize], left[:])
	copy(hashes[chainhash.HashSize:], right[:])

	newBranch := chainhash.SHA256dToHash(hashes[:])

	return &newBranch
}

// VerifyMerkleProof ...
func VerifyMerkleProof(txHash *chainhash.Hash, merkleRoot *chainhash.Hash, merkleProofs []*chainhash.Hash,
	position uint32) bool {

	hash := txHash

	for i := 0; i < len(merkleProofs); i++ {
		if ((position >> uint32(i)) & 1) == 1 {
			hash = HashMerkleBranch(merkleProofs[i], hash)
		} else {
			hash = HashMerkleBranch(hash, merkleProofs[i])
		}
	}

	if hash.Compare(merkleRoot) == 0 {
		return true
	}

	return false
}

// BuildMerkleTreeStore ...
func BuildMerkleTreeStore(txHash []*chainhash.Hash) []*chainhash.Hash {
	var nextPoT int
	txHashLen := len(txHash)

	if (txHashLen & (txHashLen - 1)) == 0 {
		nextPoT = txHashLen
	} else {
		nextPoT = 1 << (uint(math.Log2(float64(txHashLen))) + 1)
	}

	arraySize := nextPoT*2 - 1
	merkles := make([]*chainhash.Hash, arraySize)

	for i := 0; i < txHashLen; i++ {
		merkles[i] = txHash[i]
	}

	offset := nextPoT
	for i := 0; i < arraySize-1; i += 2 {
		switch {
		case merkles[i] == nil:
			merkles[offset] = nil

		case merkles[i+1] == nil:
			newHash := HashMerkleBranch(merkles[i], merkles[i])
			merkles[offset] = newHash

		default:
			newHash := HashMerkleBranch(merkles[i], merkles[i+1])
			merkles[offset] = newHash
		}
		offset++
	}
	return merkles
}

// BuildMerkleTreeRoot ...
func BuildMerkleTreeRoot(txHash []*chainhash.Hash) *chainhash.Hash {
	tree := BuildMerkleTreeStore(txHash)
	return tree[len(tree)-1]
}
