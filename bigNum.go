package rsa

import "math/big"

func mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}

func sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}

func add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}

func xor(x, y *big.Int) *big.Int {
	return big.NewInt(0).Xor(x, y)
}

func mod(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mod(x, y)
}

func modInverse(x, y *big.Int) *big.Int {
	return big.NewInt(0).ModInverse(x, y)
}
