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
	pk := calculateKey(*generator, *prime, sk)
	hasher := hashing.CreateNew(hashing.SHA256)

	return &ElGamal{signatureKey: sk, publicKey: *pk, prime: *prime, generator: *generator, hasher: hasher}
}

func (eg *ElGamal) CreateSignature(message int) (r, signature big.Int) {

	key := generateRandom()
	randomKey := calculateKey(eg.generator, eg.prime, key)
	prime := big.NewInt(SHARED_PRIME-1)

	hash := eg.hasher.Hash(eg.publicKey, []byte(strconv.Itoa(message)))
	s := big.NewInt(0).Mul(&eg.signatureKey, randomKey)
	s = s.Sub(&hash, s)
	kValue := key.Exp(&key, big.NewInt(-1), nil) 
	s = s.Mul(s,kValue)

	s = s.Mod(s, prime)

	if bigMath.Equals(s, big.NewInt(0)) {

		return eg.CreateSignature(message)
	} else {

		return *randomKey, *s
	}

}

func (eg *ElGamal) VerifySignature(message int, signature, randomKey, publicKey big.Int) bool {
	// TODO: Implement the verification signature
	zero := big.NewInt(0)
	prime := big.NewInt(SHARED_PRIME-1)

	isRandomKeyOk := bigMath.GreaterThan(&randomKey, zero) && bigMath.LessThen(&randomKey, &signature)
	isSignatureOk := bigMath.GreaterThan(&signature, zero) && bigMath.LessThen(&signature, prime)

	if !isRandomKeyOk || !isSignatureOk {

		return false
	}

	hash := eg.hasher.Hash(eg.publicKey, []byte(strconv.Itoa(message)))

	// Perform the arithmetic step in order to check that signature is valid
	generatedValue := hash.Exp(&eg.generator, &hash, nil)
	pkRandom := randomKey.Exp(&publicKey, &randomKey, nil)
	rS := randomKey.Exp(&randomKey, &signature, nil)
	pkrs := pkRandom.Mul(pkRandom, rS)


	var isSignatureValid bool = bigMath.Equals(generatedValue, pkrs)

	return isSignatureValid
}

// Generates a random integer. However not 0
func generateRandom() big.Int {
	random := rand.Int63()
	if random == 0 {
		return generateRandom()
	} else {
		return *big.NewInt(random)
	}
}

func calculateKey(base, prime, secret big.Int) *big.Int {

	result := big.NewInt(0)
	result.Exp(&base, &secret, nil)
	result.Mod(result, &prime)
	return result
}
