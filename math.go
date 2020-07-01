package rsa

import (
	"math"
	"math/big"
	"math/rand"
	"time"
)

type PrivateKey struct {
	P *big.Int
	Q *big.Int
}

type PublicKey struct {
	N *big.Int
	A *big.Int
}

var (
	minimum = 14
	maximum = 41
)

func HCF(x, y int) int {
	if y == 0 {
		return x
	}
	return HCF(y, x%y)
}

func GeneratePrime(min, max int) *big.Int {
	var x int
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(max-min) + min
	for !big.NewInt(int64(x)).ProbablyPrime(0) || HCF(x, 4) != 1 || HCF(x, 6) != 1 {
		x = rand.Intn(max-min) + min
	}
	return big.NewInt(int64(x))
}

func GeneratePrivateKey() PrivateKey {
	x := GeneratePrime(minimum, maximum)
	y := GeneratePrime(minimum, maximum)
	for y.Cmp(x) == 0 || math.Abs(float64(sub(x, y).Int64())) >= 3 {
		x = GeneratePrime(minimum, maximum)
		y = GeneratePrime(minimum, maximum)
	}
	for mul(x, y).Int64() < 200 { // otherwise the encryption gets faulty bc of too large numbers
		x = GeneratePrime(minimum, maximum)
		y = GeneratePrime(minimum, maximum)
		for y.Cmp(x) == 0 {
			y = GeneratePrime(minimum, maximum)
		}
	}
	return PrivateKey{
		P: x,
		Q: y,
	}
}

func GetPublicKey(key PrivateKey) PublicKey {
	m := int(mul(sub(key.P, big.NewInt(1)), sub(key.Q, big.NewInt(1))).Int64())
	aa := GeneratePrime(minimum, maximum)
	inv := modInverse(aa, big.NewInt(int64(m))).Int64()
	for HCF(m, int(aa.Int64())) != 1 || inv >= int64(maximum*2) { // if modInverse gets bigger than 2*the maximum for the generated numer, a will exceed the boundaries of big Integer
		aa = GeneratePrime(minimum, maximum)
		inv = modInverse(aa, big.NewInt(int64(m))).Int64()
	}
	return PublicKey{
		N: mul(key.P, key.Q),
		A: aa,
	}
}
