package commitments

import "math/big"


type CommitmentHandler[MT any] interface {
	// Takes the message and hashes it
	Commit(message MT) MT

	// Verifies that the message is equal to the commitment
	Verify(message, commitment MT, publicKey big.Int) bool
}