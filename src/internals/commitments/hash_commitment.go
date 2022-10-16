package commitments

import (
	"sec.itu.dk/ex2/internals/crypto/hashing"

	"golang.org/x/exp/slices"
)

type HashCommitment struct {
	hasher hashing.HashHandler
}

func CreateNew() *HashCommitment {

	hasher := hashing.CreateNew(hashing.SHA256)
	return &HashCommitment{hasher: hasher}
}

func (hc *HashCommitment) Commit(message, random []byte) []byte {

	return hc.hasher.Hash(random, message)
}

func (hc *HashCommitment) Verify(message, commitment, pk []byte) bool {

	hashedMessage := hc.hasher.Hash(pk, message)

	return slices.Equal(hashedMessage, commitment)
}