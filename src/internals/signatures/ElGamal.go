package signatures

type ElGamal struct {
	signatureKey int
	publicKey    int
}

func CreateNew() ElGamal {
	// TODO: Create random number in positive integers relative to a prime p-1
	// TODO: Compute the public key

	return ElGamal{signatureKey: 0, publicKey: 0}
}

// TODO: Change this to accept pointer recievers
func (eg ElGamal) CreateSignature(message int) (r int, signature int) {
	// TODO: Implement this function
	return 0, 0
}


// TODO: Change this to accept pointer recievers
func (eg ElGamal) VerifySignature(signature int, publicKey int) bool {
	// TODO: Implement the verification signature
	return false
}
