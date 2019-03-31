package blockchain

import (
	"reflect"
	"testing"

	"github.com/checksum0/go-cryptoutils/chainhash"
)

// TestHashMerkleBranch verify hashing of a merkle branch. It is the branch made out of the third
// and fourth transaction in Bitcoin Block #100,000.
// (Block #100,000, TX 3: 6359f0868171b1d194cbee1af2f16ea598ae8fad666d9b012c8ed2b79a236ec4)
// (Block #100,000, TX 4: e9a66845e05d5abc0ad04ec80f774a7e585c6e8db975962d069a522137b80c1d)
func TestHashMerkleBranch(t *testing.T) {
	tx3, _ := chainhash.NewHashFromString("6359f0868171b1d194cbee1af2f16ea598ae8fad666d9b012c8ed2b79a236ec4")
	tx4, _ := chainhash.NewHashFromString("e9a66845e05d5abc0ad04ec80f774a7e585c6e8db975962d069a522137b80c1d")
	target, _ := chainhash.NewHashFromString("8e30899078ca1813be036a073bbf80b86cdddde1c96e9e9c99e9e3782df4ae49")

	branch := HashMerkleBranch(tx3, tx4)
	if !target.IsEqual(branch) {
		t.Errorf("TestHashMerkleBranch = %s (want: %s)", branch.String(), target.String())
	}

}

// TestVerifyMerkleProof verify a merkle tree proof for the third transaction in Bitcoin block #100,000.
// (Block #100,000, TX 3: 6359f0868171b1d194cbee1af2f16ea598ae8fad666d9b012c8ed2b79a236ec4)
func TestVerifyMerkleProof(t *testing.T) {
	tx3, _ := chainhash.NewHashFromString("6359f0868171b1d194cbee1af2f16ea598ae8fad666d9b012c8ed2b79a236ec4")
	proofStr := []string{
		"e9a66845e05d5abc0ad04ec80f774a7e585c6e8db975962d069a522137b80c1d",
		"ccdafb73d8dcd0173d5d5c3c9a0770d0b3953db889dab99ef05b1907518cb815",
	}
	merkleRoot, _ := chainhash.NewHashFromString("f3e94742aca4b5ef85488dc37c06c3282295ffec960994b2c0d5ac2a25a95766")
	position := uint32(2)

	proof := make([]*chainhash.Hash, len(proofStr))

	for i := 0; i < len(proofStr); i++ {
		proof[i], _ = chainhash.NewHashFromString(proofStr[i])
	}

	verify := VerifyMerkleProof(tx3, merkleRoot, proof, position)

	if !verify {
		t.Errorf("VerifyMerkleProof = %t (want %t)", verify, true)
	}

}

// TestMerkleTreeStore builds a merkle tree store from Bitcoin Block #100,000 and verify that it
// match with the theorical merkle tree store.
// (Block #100,000: 000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506)
func TestMerkleTreeStore(t *testing.T) {
	merklesStr := []string{
		"8c14f0db3df150123e6f3dbbf30f8b955a8249b62ac1d1ff16284aefa3d06d87",
		"fff2525b8931402dd09222c50775608f75787bd2b87e56995a7bdd30f79702c4",
		"6359f0868171b1d194cbee1af2f16ea598ae8fad666d9b012c8ed2b79a236ec4",
		"e9a66845e05d5abc0ad04ec80f774a7e585c6e8db975962d069a522137b80c1d",
	}

	wantMerklesStr := []string{
		"8c14f0db3df150123e6f3dbbf30f8b955a8249b62ac1d1ff16284aefa3d06d87",
		"fff2525b8931402dd09222c50775608f75787bd2b87e56995a7bdd30f79702c4",
		"6359f0868171b1d194cbee1af2f16ea598ae8fad666d9b012c8ed2b79a236ec4",
		"e9a66845e05d5abc0ad04ec80f774a7e585c6e8db975962d069a522137b80c1d",
		"ccdafb73d8dcd0173d5d5c3c9a0770d0b3953db889dab99ef05b1907518cb815",
		"8e30899078ca1813be036a073bbf80b86cdddde1c96e9e9c99e9e3782df4ae49",
		"f3e94742aca4b5ef85488dc37c06c3282295ffec960994b2c0d5ac2a25a95766",
	}

	merkles := make([]*chainhash.Hash, len(merklesStr))
	wantMerkles := make([]*chainhash.Hash, len(wantMerklesStr))

	for i := 0; i < len(merklesStr); i++ {
		merkles[i], _ = chainhash.NewHashFromString(merklesStr[i])
	}

	for i := 0; i < len(wantMerklesStr); i++ {
		wantMerkles[i], _ = chainhash.NewHashFromString(wantMerklesStr[i])
	}

	merkleTree := BuildMerkleTreeStore(merkles)

	if !reflect.DeepEqual(merkleTree, wantMerkles) {
		t.Errorf("BuildMerkleTreeStore = %v (want %v)", merkleTree, wantMerkles)
	}
}

// TestMerkleTreeRoot build the merkle tree root from Bitcoin block #100,000 and verify that it
// matches with the real merkle tree root.
// (Block #100,000: 000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506)
func TestMerkleTreeRoot(t *testing.T) {
	merklesStr := []string{
		"8c14f0db3df150123e6f3dbbf30f8b955a8249b62ac1d1ff16284aefa3d06d87",
		"fff2525b8931402dd09222c50775608f75787bd2b87e56995a7bdd30f79702c4",
		"6359f0868171b1d194cbee1af2f16ea598ae8fad666d9b012c8ed2b79a236ec4",
		"e9a66845e05d5abc0ad04ec80f774a7e585c6e8db975962d069a522137b80c1d",
	}

	merkles := make([]*chainhash.Hash, len(merklesStr))

	for i := 0; i < len(merklesStr); i++ {
		merkles[i], _ = chainhash.NewHashFromString(merklesStr[i])
	}

	merkleRoot := BuildMerkleTreeRoot(merkles)
	wantRoot, _ := chainhash.NewHashFromString("f3e94742aca4b5ef85488dc37c06c3282295ffec960994b2c0d5ac2a25a95766")

	if !reflect.DeepEqual(merkleRoot, wantRoot) {
		t.Errorf("BuildMerkleTreeRoot = got %v (want %v)", merkleRoot, wantRoot)
	}
}
