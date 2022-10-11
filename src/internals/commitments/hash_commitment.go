package commitments

import (
	"math/big"

	"sec.itu.dk/ex2/internals/crypto/hashing"
)

type HashCommitment struct {
	hasher hashing.HashHandler
}

func CreateNew() *HashCommitment {

	hasher := hashing.CreateNew(hashing.SHA256)
	return &HashCommitment{hasher: hasher}
}

func (hc *HashCommitment) Commit(message, pk big.Int) big.Int {

	return hc.hasher.Hash(pk, message.Bytes())
}

func (hc *HashCommitment) Verify(message, commitment, pk big.Int) bool {

	hashedMessage := hc.hasher.Hash(pk, message.Bytes())

	return hashedMessage.Cmp(&commitment) == 0
}