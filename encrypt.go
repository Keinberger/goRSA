package rsa

import (
	"math/big"
)

func Encrypt(msg int, pubKey PublicKey) int64 {
	msgBig := big.NewInt(int64(msg))
	aBig := big.NewInt(int64(pubKey.A))
	xA := xor(msgBig, aBig)
	res := mod(xA, big.NewInt(int64(pubKey.N)))
	return res.Int64()
}

func EncryptBytes(msg []byte, pubKey PublicKey) []byte {
	encrypted := make([]byte, len(msg))
	for k, v := range msg {
		encr := Encrypt(int(v), pubKey)
		encrypted[k] = byte(encr)
	}
	return encrypted
}
