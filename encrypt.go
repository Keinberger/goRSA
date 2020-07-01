package rsa

import "math/big"

func Encrypt(msg int64, pubKey PublicKey) int64 {
	msgBig := big.NewInt(msg)
	aBig := pubKey.A
	xA := xor(msgBig, aBig)
	res := mod(xA, pubKey.N)
	return res.Int64()
}

func EncryptBytes(msg []byte, pubKey PublicKey) []byte {
	encrypted := make([]byte, len(msg))
	for k, v := range msg {
		encr := Encrypt(int64(v), pubKey)
		encrypted[k] = byte(encr)
	}
	return encrypted
}
