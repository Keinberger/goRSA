package rsa

import "math/big"

// Encrypt encodes a msg of type int64 and returns the encrypted msg
func Encrypt(msg int64, pubKey PublicKey) int64 {
	msgBig := big.NewInt(msg)
	aBig := pubKey.A
	xA := xor(msgBig, aBig)
	res := mod(xA, pubKey.N)
	return res.Int64()
}

// DecryptBytes encodes a byte array by converting the bytes into int64 and passing the number through Encrypt()
// afterwards it will convert the numbers into bytes and put these encoded numbers back into a byte array and return it
func EncryptBytes(msg []byte, pubKey PublicKey) []byte {
	encrypted := make([]byte, len(msg))
	for k, v := range msg {
		encr := Encrypt(int64(v), pubKey)
		encrypted[k] = byte(encr)
	}
	return encrypted
}
