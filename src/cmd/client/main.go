package main

import (
	"fmt"

	"sec.itu.dk/ex2/internals/signatures"
)

func main() {
	var signatureHandler signatures.SignatureHandler[int] = signatures.CreateNew()
	var _, _ = signatureHandler.CreateSignature(2)
	fmt.Println("Hello World!")
}
