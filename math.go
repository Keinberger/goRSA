package rsa

import (
	"math/rand"
	"time"
)

type PrivateKey struct {
	P int
	Q int
}

type PublicKey struct {
	N int
	A int
}

func HCF(x, y int) int {
	if y == 0 {
		return x
	}
	return HCF(y, x%y)
}

func GeneratePrime(min, max int) int {
	var x int
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(max-min+1) + min
	for HCF(x, 2) != 1 || HCF(x, 3) != 1 || HCF(x, 4) != 1 || HCF(x, 5) != 1 || HCF(x, 7) != 1 || x == 1 || x == 0 {
		x = rand.Intn(max-min+1) + min
	}
	return x
}

func GeneratePrivateKey() PrivateKey {
	x := GeneratePrime(1, 20)
	y := GeneratePrime(1, 20)
	for x == y {
		y = GeneratePrime(1, 20)
	}
	for x*y > 144 { // otherwise the encryption gets faulty
		x = GeneratePrime(1, 20)
		y = GeneratePrime(1, 20)
		for x == y {
			y = GeneratePrime(1, 20)
		}
	}
	return PrivateKey{
		P: x,
		Q: y,
	}
}

func GetPublicKey(key PrivateKey) PublicKey {
	z := GeneratePrime(1, 12)
	for (key.P-1)*(key.Q-1)%z == 0 {
		z = GeneratePrime(1, 12)
	}
	return PublicKey{
		N: key.P * key.Q,
		A: z,
	}
}
