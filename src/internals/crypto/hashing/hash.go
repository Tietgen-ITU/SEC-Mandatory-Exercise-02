package hashing

import (
	"hash"
	"math/big"

	"golang.org/x/crypto/sha3"
)

type HashType int32

const (
	SHA256 HashType = 256
	SHA384 HashType = 384
	SHA512 HashType = 512
)

type HashHandler interface {
	// Hashes message with the given integer key
	Hash(key big.Int, message []byte) big.Int

	// Compares the message with the other integer hash and integer key 
	Compare(message []byte, key, hashValue big.Int) bool
}

type Hasher struct {
	internalHasher hash.Hash
} 

func CreateNew(hashType HashType) *Hasher {

	var hash hash.Hash

	switch hashType {
	case SHA256:
		hash = sha3.New256()		
	case SHA384:
		hash = sha3.New384()
	case SHA512:
		hash = sha3.New512()
	}
	return &Hasher{internalHasher: hash}
}

func (hasher *Hasher) Hash(key big.Int, message []byte) big.Int {

	// Reset the hasher after the hash value has been created
	defer hasher.internalHasher.Reset()
	
	// Provide the key and value to be hashed
	hasher.internalHasher.Write(key.Bytes())
	hasher.internalHasher.Write(message)

	// Get the resulting size of the byte array
	size := hasher.internalHasher.Size()
	resultBytes := hasher.internalHasher.Sum(make([]byte, size))

	resultInt := big.NewInt(0).SetBytes(resultBytes)
	return *resultInt 
}

func (hasher *Hasher) Compare(message []byte, key, hashValue big.Int) bool {

	var hashedMessage big.Int = hasher.Hash(key, message)

	return hashValue.Cmp(&hashedMessage) == 0 // If the value is zero then it means that they are the same
}