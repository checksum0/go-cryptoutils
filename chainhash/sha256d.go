package chainhash

import "crypto/sha256"

// SHA256ToBytes ...
func SHA256ToBytes(b []byte) []byte {
	once := sha256.Sum256(b)

	return once[:]
}

// SHA256ToHash ...
func SHA256ToHash(b []byte) Hash {
	once := Hash(sha256.Sum256(b))

	return once
}

// SHA256dToBytes ...
func SHA256dToBytes(b []byte) []byte {
	once := sha256.Sum256(b)
	twice := sha256.Sum256(once[:])

	return twice[:]
}

// SHA256dToHash ...
func SHA256dToHash(b []byte) Hash {
	once := sha256.Sum256(b)
	twice := Hash(sha256.Sum256(once[:]))

	return twice
}
