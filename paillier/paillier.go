package paillier

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

const KEY_LENGTH = 3072

func GenerateKey(random io.Reader) (*PrivateKey, error) {
	primeSize := KEY_LENGTH / 2

	p, err := rand.Prime(random, primeSize)
	if err != nil {
		panic(err)
	}

	q, err := rand.Prime(random, primeSize)
	if err != nil {
		panic(err)
	}

	n := new(big.Int).Mul(p, q)

	return &PrivateKey{
		PublicKey: PublicKey{
			N:  n,
			NN: new(big.Int).Mul(n, n),
			G:  new(big.Int).Add(n, new(big.Int).SetInt64(1)),
		},
		p: p,
		q: q,
		n: n,
	}, nil

}

type PrivateKey struct {
	PublicKey
	p *big.Int
	q *big.Int
	n *big.Int
}

type PublicKey struct {
	N  *big.Int
	NN *big.Int
	G  *big.Int
}

func Encrypt(pubKey *PublicKey, plainText []byte) ([]byte, error) {
	r, err := rand.Int(rand.Reader, pubKey.N)
	if err != nil {
		return nil, err
	}

	m := new(big.Int).SetBytes(plainText)
	if pubKey.N.Cmp(m) < 1 {
		return nil, fmt.Errorf("message too long (can't encrypt message larger than the public key)")
	}

	// c = g^m * r^n mod n^2
	left := new(big.Int).Exp(pubKey.G, m, pubKey.NN)
	right := new(big.Int).Exp(r, pubKey.N, pubKey.NN)
	c := new(big.Int).Mul(left, right)
	c = new(big.Int).Mod(c, pubKey.NN)

	return c.Bytes(), nil
}

func Decrypt(privKey *PrivateKey, cipherText []byte) ([]byte, error) {
	c := new(big.Int).SetBytes(cipherText)
	if privKey.PublicKey.NN.Cmp(c) < 1 {
		return nil, fmt.Errorf("message too long (can't decrypt message larger than the public key)")
	}

	// m = l ( c^λ mod n^2) * μ mod n
	pMinusOne := new(big.Int).Sub(privKey.p, new(big.Int).SetInt64(1))
	qMinusOne := new(big.Int).Sub(privKey.q, new(big.Int).SetInt64(1))

	λ := new(big.Int).Mul(pMinusOne, qMinusOne)
	x := new(big.Int).Exp(c, λ, privKey.NN)
	lx := l(x, privKey.n)

	μ := new(big.Int).ModInverse(λ, privKey.N)

	m := new(big.Int).Mul(lx, μ)
	m.Mod(m, privKey.n)
	return m.Bytes(), nil
}

func l(u *big.Int, n *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Sub(u, new(big.Int).SetInt64(1)), n)
}

func Add(publicKey *PublicKey, cipher, plain []byte) ([]byte, error) {
	// m_1 + m_2 mod n = dec ( enc(m_1,r_1) * enc(m_2, r_2) mod n^2 )
	m2B, err := Encrypt(publicKey, plain)
	if err != nil {
		return nil, err
	}

	m1 := new(big.Int).SetBytes(cipher)
	m2 := new(big.Int).SetBytes(m2B)
	additionRes := new(big.Int).Mul(m1, m2)
	res := new(big.Int).Mod(additionRes, publicKey.NN)
	return res.Bytes(), nil
}

func Mul(publicKey *PublicKey, cipher, plain []byte) []byte {
	m1 := new(big.Int).SetBytes(cipher)
	m2 := new(big.Int).SetBytes(plain)
	res := new(big.Int).Exp(m1, m2, publicKey.NN)

	return res.Bytes()
}
