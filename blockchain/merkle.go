package blockchain

import (
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
