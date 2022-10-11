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

func (eg *ElGamal) PublicKey() big.Int {

	return eg.publicKey
}

func (eg *ElGamal) Sign(message int) (r, signature big.Int) {

	primeMinusOne := big.NewInt(SHARED_PRIME-1)

	// Get k relatively prime to SHARED_PRIME - 1
	key := generateRandomRelativelyPrime(primeMinusOne)

	// Calculate random key a.k.a. r
	randomKey := calculateKey(eg.generator, eg.prime, key)

	// Calculate s
	hash := eg.hasher.Hash(eg.publicKey, []byte(strconv.Itoa(message)))
	sAux1 := new(big.Int).Sub(&hash, big.NewInt(0).Mul(&eg.signatureKey, randomKey))
	kMulInverse := new(big.Int).ModInverse(&key, &eg.prime) 
	result := new(big.Int).Mul(sAux1,kMulInverse)

	s := new(big.Int).Mod(result, primeMinusOne)

	if bigMath.Equals(s, big.NewInt(0)) {

		return eg.Sign(message)
	} else {

		return *randomKey, *s
	}
}

func (eg *ElGamal) Verify(message int, signature, randomKey, publicKey big.Int) bool {

	zero := big.NewInt(0)
	primeMinusOne := big.NewInt(SHARED_PRIME-1)

	isRandomKeyOk := bigMath.GreaterThan(&randomKey, zero) && bigMath.LessThen(&randomKey, &eg.prime)
	isSignatureOk := bigMath.GreaterThan(&signature, zero) && bigMath.LessThen(&signature, primeMinusOne)

	if !isRandomKeyOk || !isSignatureOk {

		return false
	}

	hash := eg.hasher.Hash(publicKey, []byte(strconv.Itoa(message)))

	// Perform the arithmetic step in order to check that signature is valid
	rS := new(big.Int).Exp(&randomKey, &signature, &eg.prime)
	pkRandom := new(big.Int).Exp(&publicKey, &randomKey, &eg.prime)
	pkrs := new(big.Int).Mul(pkRandom, rS)
	pkrs.Mod(pkrs, &eg.prime)
	generatedValue := new(big.Int).Exp(&eg.generator, &hash, &eg.prime)

	var isSignatureValid bool = bigMath.Equals(generatedValue, pkrs)

	return isSignatureValid
}

// Generates a random integer for the set Z^*_p. This means the set of intergers between 1, 2.. p-1
func generateRandom() big.Int {
	random := rand.Int63n(SHARED_PRIME)
	if random <= 0 {
		return generateRandom()
	} else {
		return *big.NewInt(random)
	}
}


func generateRandomRelativelyPrime(p *big.Int) big.Int {

	one := big.NewInt(1)
	random := generateRandom()
	gcd := new(big.Int).GCD(nil, nil, &random, p)

	if bigMath.Equals(gcd, one)  {
		return generateRandomRelativelyPrime(p)
	} else {
		return random
	}
}

func calculateKey(base, prime, secret big.Int) *big.Int {

	result := big.NewInt(0)
	result.Exp(&base, &secret, &prime)
	return result
}
