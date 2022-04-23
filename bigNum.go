package rsa

import "math/big"

// mul() multiples big numbers x and y and returns the result
func mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}

// sub() subtracts y off x and returns the result
func sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}

// xor() returns x to the power of y
func xor(x, y *big.Int) *big.Int {
	return big.NewInt(0).Xor(x, y)
}

// mod() returns x%y
func mod(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mod(x, y)
}

// modInverse() calculates the mod inverse of x and y and returns the result
func modInverse(x, y *big.Int) *big.Int {
	return big.NewInt(0).ModInverse(x, y)
}
