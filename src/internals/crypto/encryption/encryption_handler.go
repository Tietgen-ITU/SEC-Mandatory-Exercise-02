package encryption

import "math/big"

type EncryptionHandler[MT any] interface {
	Encrypt(message MT, pk big.Int) big.Int
	Decrypt(cipher big.Int, pk big.Int) MT
}