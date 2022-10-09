package signatures

type SignatureHandler[MT any] interface {
	CreateSignature(MT) (int, int)
	VerifySignature(int, int) bool
}