package signatures

import "math/big"

type SignatureHandler[MT any] interface {

	// Creates a signature based on the message
	Sign(MT) (random, s big.Int)

	// Verifies the signature
	Verify(message int, s, random, pk big.Int) bool

	// Gets the public key used for the signature handler
	PublicKey() big.Int
}
