package commitments


type CommitmentHandler[MT any] interface {
	// Takes the message and hashes it
	Commit(MT) MT
	Verify(MT, publicKey int) bool
}