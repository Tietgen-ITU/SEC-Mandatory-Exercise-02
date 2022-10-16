package commitments

import (
	"math/big"
	"testing"
)

func TestCommitmentAndVerifyWithSamePublicKeyReturnTrue(t *testing.T) {

	// Setup test
	commitmentHandler := CreateNew()
	message := big.NewInt(425).Bytes()
	publicKey := big.NewInt(6669).Bytes()
	
	// Act
	commitment := commitmentHandler.Commit(message, publicKey)
	resultOk := commitmentHandler.Verify(message, commitment, publicKey)

	// Assert
	if !resultOk {
		t.FailNow()
	}
}

func TestCommitmentAndVerifyWithDifferentPublicKeyReturnFalse(t *testing.T) {


	// Setup test
	commitmentHandler := CreateNew()
	message := big.NewInt(425).Bytes()
	publicKey := big.NewInt(6669).Bytes()
	falsePublicKey := big.NewInt(45).Bytes()
	
	// Act
	commitment := commitmentHandler.Commit(message, publicKey)
	resultOk := commitmentHandler.Verify(message, commitment, falsePublicKey)

	// Assert
	if resultOk {
		t.FailNow()
	}
}