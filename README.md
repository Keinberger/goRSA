# goRSA

## Usage

Import the library like that: `import rsa "github.com/keinberger/goRSA"`

### Generating a key pair

The package contains two separe functions for generating the privateKey and publicKey. Both will return variables of type privateKey or publicKey.

One has to generate the privateKey before the publicKey, because the publicKey dependends on the privateKey.

E.g.: 
```
privateKey := rsa.GeneratePrivateKey()
publicKey := rsa.GeneratePublicKey(privateKey)
```

### Using the keys to encode/decode numbers

After having generated a key pair containing of privateKey and publicKey, you may use these variables to encrypt and decrypt numbers of type int64.

The package contains one function each for encrypting and decryping, both will return a variable of type int64.

Encrypting a number works like that: 
```
privateKey := rsa.GeneratePrivateKey()
publicKey := rsa.GetPublicKey(privateKey)

var numberToEncrypt int64 = 100
encryptedNumber := rsa.Encrypt(numberToEncrypt, publicKey)
```

Decrypting the encrypted number works like that: 
```
decryptedNumber := rsa.Decrypt(encryptedNumber, privateKey, publicKey)
```

The variables `numberToEncrypt` and `decryptedNumber`, used in this example, should be the same.

### Encoding/Decoding strings

The package also supports the encryption/decryption of byte-slices (strings).

Here's how to use that:

Encrypting:```
privateKey := rsa.GeneratePrivateKey()
publicKey := rsa.GetPublicKey(privateKey)

stringToEncrypt := "I love RSA"
encryptedString := string(rsa.EncryptBytes([]byte(stringToEncrypt), publicKey))
```

Decrypting:```
decryptedString := string(rsa.DecryptBytes([]byte(encryptedString), privateKey, publicKey))
```

The variables `stringToEncrypt` and `decryptedString`, used in this example, should be the same.

## Important

There might still be some bugs in the program.
