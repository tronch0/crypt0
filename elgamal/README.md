# ElGamal Cryptosystem

## Explanation
The ElGamal cryptosystem is an asymmetric algorithm for public-key cryptography based on the discrete logarithm problem. It was invented by Taher Elgamal in 1985.

This package contains an implementation of the ElGamal cryptosystem in Go. It supports the following operations:

* Key pair generation
* Encryption of messages
* Decryption of ciphertexts
* Homomorphic multiplication of ciphertexts

## Usage

### Public Parameters Generation
```go
// Generate public parameters for ElGamal key generation
param, err := generatePublicParam(256)
if err != nil {
panic(err)
}
```

### Key Gen
```go
// Generate a new key pair with a given prime modulus and generator
p := new(big.Int).SetBytes([]byte{...})
g := new(big.Int).SetBytes([]byte{...})
privKey, err := GenerateKeyPair(p, g)
if err != nil {
panic(err)
}
```
### Encrypt
```go
// Encrypt a message using the public key
pubKey := privKey.PublicKey
m := new(big.Int).SetBytes([]byte{...})
c1, c2, err := Encrypt(&pubKey, m)
if err != nil {
panic(err)
}
```
### Decrypt
```go
// Decrypt a ciphertext using the private key
m, err := Decrypt(privKey, c1, c2)
if err != nil {
panic(err)
}
```
### Homomorphic Multiplication
```go
// Multiply two ciphertexts homomorphically
pubKey := privKey.PublicKey
c1a, c2a, err := Encrypt(&pubKey, new(big.Int).SetInt64(10))
if err != nil {
    panic(err)
}
c1b, c2b, err := Encrypt(&pubKey, new(big.Int).SetInt64(5))
if err != nil {
    panic(err)
}
c1, c2, err := HomomorphicMul(&pubKey, c1a, c2a, c1b, c2b)
if err != nil {
    panic(err)
}
```

## Warning
This library was created primarily for education purposes, with future application for a course project. You should **NOT USE THIS CODE IN PRODUCTION SYSTEMS**.
