package math

import "math/big"

func GreaterThan(a *big.Int, b *big.Int) bool {

	return a.Cmp(b)	== 1
}

func LessThen(a, b *big.Int) bool {
	return a.Cmp(b) == -1
}

func Equals(a, b *big.Int) bool {

	return a.Cmp(b) == 0
}

func GreaterThanOrEquals(a, b *big.Int) bool {

	return a.Cmp(b) >= 0
}

func LessThanOrEquals(a, b *big.Int) bool {

	return a.Cmp(b) <= 0
}