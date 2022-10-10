package signatures

import "math/big"

type SignatureHandler[MT any] interface {
	CreateSignature(MT) (random, s big.Int)
	VerifySignature(message int, s, random, pk big.Int) bool
}
