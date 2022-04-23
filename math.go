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

const (
	minimum = 14
	maximum = 41
)

// hcf calculates the highest common factor of x and y
func hcf(x, y int) int {
	if y == 0 {
		return x
	}
	return hcf(y, x%y)
}

// GeneratePrime generates a prime number of type big.Int between min and max
func GeneratePrime(min, max int) *big.Int {
	var x int
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(max-min) + min
	for !big.NewInt(int64(x)).ProbablyPrime(0) || hcf(x, 4) != 1 || hcf(x, 6) != 1 {
		x = rand.Intn(max-min) + min
	}
	return big.NewInt(int64(x))
}

// GeneratePrivateKey returns a newly generated private Key containing two prime numbers
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

// GetPublicKey returns the public Key associated to the given privateKey containing the multiplication of the two numbers inside of the private Key
// and a newly generated prime number A, being coprime to the multiplication of the two numbers inside of the privateKey, each minus 1 ( (p-1)*(q-1) )
func GetPublicKey(key PrivateKey) (pubKey PublicKey) {
	pubKey = PublicKey{
		A: GeneratePrime(minimum, maximum),
		N: mul(key.P, key.Q),
	}

	m := int(mul(sub(key.P, big.NewInt(1)), sub(key.Q, big.NewInt(1))).Int64())
	inv := modInverse(pubKey.A, big.NewInt(int64(m))).Int64()

	for hcf(m, int(pubKey.A.Int64())) != 1 || inv >= int64(maximum*2) { // if modInverse gets bigger than 2*the maximum for the generated numer, a will exceed the boundaries of big Integer
		pubKey.A = GeneratePrime(minimum, maximum)
		inv = modInverse(pubKey.A, big.NewInt(int64(m))).Int64()
	}

	return
}
