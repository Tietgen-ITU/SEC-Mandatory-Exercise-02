package crypto

import (
	"math/big"
	"testing"
)

func TestMessageSignatureVerificationSucceed(t *testing.T) {

	handler := CreateNew()
	msg := 5

	r, s := handler.Sign(msg)
	if ok := handler.Verify(msg, s, r, handler.PublicKey()); !ok {
		t.Fatalf("Handler did not return ok after signature verification")
	} 
}

func TestMessageSignatureWithWrongKeyVerificationFailed(t *testing.T) {

	handler := CreateNew()
	msg := 5

	r, s := handler.Sign(msg)
	if ok := handler.Verify(msg, s, r, *big.NewInt(500092230000)); ok {
		t.Fatalf("Handler returned ok after wrong key inserted")
	} 
}

func TestEncryptAndDecryptMessageWithCorrectKeysReturnCorrectMessage(t *testing.T) {

	t.Fatalf("Not implemented")
}

func TestEncryptAndDecryptMessageWithWrongKeyReturnsIncorrectMessage(t *testing.T) {

	t.Fatalf("Not implemented")
}