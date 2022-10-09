package signatures

import (
	"math/big"
	"math/rand"
	"strconv"

	"sec.itu.dk/ex2/internals/crypto/hashing"
	bigMath "sec.itu.dk/ex2/internals/math"
)

const (
	SHARED_GENERATOR int64 = 666
	SHARED_PRIME     int64 = 6661
)

type ElGamal struct {
	signatureKey big.Int
	publicKey    big.Int
	prime        big.Int
	generator    big.Int
	hasher       hashing.HashHandler
}

// Generates a new ElGamal signature scheme with keys
func CreateNew() *ElGamal {

	generator := big.NewInt(SHARED_GENERATOR)
	prime := big.NewInt(SHARED_PRIME)

	sk := generateRandom()
	secret := big.NewInt(sk)
	pk := calculateKey(*generator, *prime, *secret)
	hasher := hashing.CreateNew(hashing.SHA256)

	return &ElGamal{signatureKey: *secret, publicKey: *pk, prime: *prime, generator: *generator, hasher: hasher}
}

func (eg *ElGamal) CreateSignature(message int) (r, signature big.Int) {

	key := generateRandom()
	randomKey := calculateKey(eg.generator, eg.prime, *big.NewInt(key))

	hash := eg.hasher.Hash(eg.publicKey, []byte(strconv.Itoa(message)))
	s := 0 // TODO: Implement the generation of the signature
	return *randomKey, *s
}

func (eg *ElGamal) VerifySignature(signature, randomKey, publicKey big.Int) bool {
	// TODO: Implement the verification signature
	zero := big.NewInt(0)
	prime := big.NewInt(SHARED_PRIME-1)

	isRandomKeyOk := bigMath.GreaterThan(&randomKey, zero) && bigMath.LessThen(&randomKey, &signature)
	isSignatureOk := bigMath.GreaterThan(&signature, zero) && bigMath.LessThen(&signature, prime)

	if !isRandomKeyOk || !isSignatureOk {

		return false
	}

	// hash := eg.hasher.Hash()

	// TODO: Implement the signature verification calculation
	var isSignatureValid bool = false

	return isSignatureValid
}

// Generates a random integer. However not 0
func generateRandom() int64 {
	random := rand.Int63()
	if random == 0 {
		return 1
	} else {
		return random
	}
}

func calculateKey(base, prime, secret big.Int) *big.Int {

	result := big.NewInt(0)
	result.Exp(&base, &secret, nil)
	result.Mod(result, &prime)
	return result
}
