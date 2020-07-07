package rsa

import (
	"math/big"
)

// Decrypt decodes a msg of type int64 and returns the decrypted number int64
func Decrypt(msg int64, priKey PrivateKey, pubKey PublicKey) int64 {
	mBig := mul(sub(priKey.P, big.NewInt(1)), sub(priKey.Q, big.NewInt(1)))
	aBig := pubKey.A
	nBig := pubKey.N

	b := modInverse(aBig, mBig)
	msgBig := big.NewInt(msg)

	res := mod(xor(msgBig, b), nBig)
	return res.Int64()
}

// DecryptBytes decodes a byte array by converting the bytes into int64 and passing the number through Decrypt()
// afterwards it will convert the numbers into bytes and put these decoded numbers back into a byte array and return it
func DecryptBytes(msg []byte, priKey PrivateKey, pubKey PublicKey) []byte {
	decrypted := make([]byte, len(msg))
	for k, v := range msg {
		res := Decrypt(int64(v), priKey, pubKey)
		decrypted[k] = byte(res)
	}
	return decrypted
}
