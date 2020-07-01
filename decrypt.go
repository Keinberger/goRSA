package rsa

import (
	"math/big"
)

func Decrypt(msg int64, priKey PrivateKey, pubKey PublicKey) int64 {
	m := (priKey.P - 1) * (priKey.Q - 1)
	aBig := big.NewInt(int64(pubKey.A))
	mBig := big.NewInt(int64(m))
	nBig := big.NewInt(int64(pubKey.N))

	b := modInverse(aBig, mBig)
	msgBig := big.NewInt(msg)

	res := mod(xor(msgBig, b), nBig)
	return res.Int64()
}

func DecryptBytes(msg []byte, priKey PrivateKey, pubKey PublicKey) []byte {
	decrypted := make([]byte, len(msg))
	for k, v := range msg {
		res := Decrypt(int64(v), priKey, pubKey)
		decrypted[k] = byte(res)
	}
	return decrypted
}
