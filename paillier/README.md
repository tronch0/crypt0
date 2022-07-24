# Paillier cryptosystem

## Explanation
The Paillier cryptosystem, invented by and named after Pascal Paillier in 1999, is a probabilistic asymmetric algorithm for public key cryptography.

The implementation is based on the [Public-Key Cryptosystems Based on Composite Degree Residuosity Classes](https://link.springer.com/content/pdf/10.1007%2F3-540-48910-X_16.pdf) papper.

The package support the following operations
* Encrypt integers
* Decrypt
* Encrypted integers can be added to unencrypted integers
* Encrypted integers can be multiplied by an unencrypted integer

## Use the library
### Key Gen
```go
// Generate a 3072-bit private key.
privKey, err := GenerateKey(rand.Reader)
if err != nil {
    panic(err)
}
```
### Encrypt
```go
n := new(big.Int).SetInt64(15)

cipher, err := Encrypt(&keys.PublicKey, n.Bytes())
if err != nil {
    panic(err)
}
```
### Decrypt
```go
res, err := Decrypt(keys, cipher)
if err != nil {
    panic(err)
}
```
### Addition
```go
toAdd := new(big.Int).SetInt64(10).Bytes()

addResEnc, err := Add(&keys.PublicKey, cipher, toAdd)
if err != nil {
    panic(err)
}
```
### Multiplication
```go
toMul := new(big.Int).SetInt64(10).Bytes()

mulResEnc, err := Mul(&keys.PublicKey, cipher, toMul)
if err != nil {
    panic(err)
}
```

## Warning
This library was created primarily for education purposes, with future application for a course project. You should **NOT USE THIS CODE IN PRODUCTION SYSTEMS**.

[papper](https://link.springer.com/content/pdf/10.1007%2F3-540-48910-X_16.pdf)

