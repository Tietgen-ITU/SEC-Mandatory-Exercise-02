package commitments

import (
	"math/big"
	"testing"
)

func TestCommitmentAndVerifyWithSamePublicKeyReturnTrue(t *testing.T) {

	// Setup test
	commitmentHandler := CreateNew()
	message := big.NewInt(425)
	publicKey := big.NewInt(6669)
	
	// Act
	commitment := commitmentHandler.Commit(*message, *publicKey)
	resultOk := commitmentHandler.Verify(*message, commitment, *publicKey)

	// Assert
	if !resultOk {
		t.FailNow()
	}
}

func TestCommitmentAndVerifyWithDifferentPublicKeyReturnFalse(t *testing.T) {


	// Setup test
	commitmentHandler := CreateNew()
	message := big.NewInt(425)
	publicKey := big.NewInt(6669)
	falsePublicKey := big.NewInt(45)
	
	// Act
	commitment := commitmentHandler.Commit(*message, *publicKey)
	resultOk := commitmentHandler.Verify(*message, commitment, *falsePublicKey)

	// Assert
	if resultOk {
		t.FailNow()
	}
}