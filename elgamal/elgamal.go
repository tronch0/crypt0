package elgamal

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type PublicKey struct {
	P *big.Int // prime modulus
	G *big.Int // generator
	H *big.Int // public key
}

type PrivateKey struct {
	PublicKey
	X *big.Int // private key
}

// GenerateKeyPair generates a new Exponential ElGamal key pair.
func GenerateKeyPair(p *big.Int, g *big.Int) (*PrivateKey, error) {
	priv := new(PrivateKey)
	priv.P = p
	priv.G = g

	// Choose random private key x, 0 <= x < p-1
	x, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(1)))
	if err != nil {
		return nil, fmt.Errorf("error generating private key: %v", err)
	}
	priv.X = x

	// Compute public key h = g^x mod p
	priv.H = new(big.Int).Exp(g, x, p)

	return priv, nil
}

// Encrypt encrypts a message using the given public key.
func Encrypt(pub *PublicKey, m *big.Int) (*big.Int, *big.Int, error) {
	// Choose random r, 0 <= r < p-1
	r, err := rand.Int(rand.Reader, new(big.Int).Sub(pub.P, big.NewInt(1)))
	if err != nil {
		return nil, nil, fmt.Errorf("error generating random number: %v", err)
	}

	// Compute c1 = g^r mod p
	c1 := new(big.Int).Exp(pub.G, r, pub.P)

	// Compute c2 = m * h^r mod p
	c2 := new(big.Int).Mul(m, new(big.Int).Exp(pub.H, r, pub.P))
	c2.Mod(c2, pub.P)

	return c1, c2, nil
}

// Decrypt decrypts a ciphertext using the given private key.
func Decrypt(priv *PrivateKey, c1 *big.Int, c2 *big.Int) (*big.Int, error) {
	// Compute s = c1^x mod p
	s := new(big.Int).Exp(c1, priv.X, priv.P)

	// Compute m = c2 * s^-1 mod p
	s.ModInverse(s, priv.P)
	m := new(big.Int).Mul(c2, s)
	m.Mod(m, priv.P)

	return m, nil
}

// HomomorphicMul adds two ciphertexts together homomorphically.
func HomomorphicMul(pub *PublicKey, c1a *big.Int, c2a *big.Int, c1b *big.Int, c2b *big.Int) (*big.Int, *big.Int, error) {
	// Compute c1 = c1a * c1b mod p
	c1 := new(big.Int).Mul(c1a, c1b)
	c1.Mod(c1, pub.P)

	// Compute c2 = c2a * c2b mod p
	c2 := new(big.Int).Mul(c2a, c2b)
	c2.Mod(c2, pub.P)

	return c1, c2, nil
}

type PublicParam struct {
	G *big.Int // generator
	Q *big.Int
	P *big.Int // prime
}

func generatePublicParam(bits int) (*PublicParam, error) {
	q, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	two := new(big.Int).SetInt64(2)
	p := new(big.Int).Add(new(big.Int).Mul(q, two), new(big.Int).SetInt64(1))

	for !p.ProbablyPrime(20) {
		q, err = rand.Prime(rand.Reader, bits)
		if err != nil {
			return nil, err
		}

		p = new(big.Int).Add(new(big.Int).Mul(q, two), new(big.Int).SetInt64(1))
	}

	h, err := rand.Int(rand.Reader, p)
	if err != nil {
		return nil, err
	}

	g := new(big.Int).Exp(h, two, p)

	return &PublicParam{G: g, P: p, Q: q}, nil
}
