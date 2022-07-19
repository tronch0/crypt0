package shamir

import (
	"github.com/tronch0/crypt0/bigint"
	"github.com/tronch0/crypt0/field"
	"math/big"
)

type Point struct {
	X *field.Element
	Y *field.Element
}

func Split(secret *field.Element, numOfShares, threshold int) ([]Point, error) {
	fieldOrder := secret.GetOrder()

	coefs := make([]*field.Element, threshold)
	coefs[0] = secret

	for i := 1; i < len(coefs); i++ {
		coefs[i] = field.New(bigint.GetRandom(), fieldOrder)
	}

	shares := make([]Point, numOfShares)

	for i := 0; i < numOfShares; i++ {
		x := field.New(new(big.Int).SetInt64(int64(i+1)), fieldOrder)
		shares[i] = Point{X: x, Y: computePolynomial(coefs, x)}
	}

	return shares, nil
}

func computePolynomial(coefs []*field.Element, x *field.Element) *field.Element {
	res := coefs[0]
	for i := 1; i < len(coefs); i++ {

		expoRes := x.Expo(new(big.Int).SetInt64(int64(i)))
		mulRes := coefs[i].Mul(expoRes)

		res = res.Add(mulRes)
	}

	return res
}

func Recover(shares []Point) *field.Element {
	fieldOrder := shares[0].X.GetOrder()
	res := field.New(new(big.Int).SetInt64(0), fieldOrder)

	for i := 0; i < len(shares); i++ {
		ci := field.New(bigint.GetRandom(), fieldOrder)
		numerator := field.New(new(big.Int).SetInt64(1), fieldOrder)
		denominator := field.New(new(big.Int).SetInt64(1), fieldOrder)

		for j := 0; j < len(shares); j++ {

			if i == j {
				continue
			}

			toMulWithNum := field.New(new(big.Int).SetInt64(0), fieldOrder).Sub(shares[j].X)
			toMulWithNDen := shares[i].X.Sub(shares[j].X)

			numerator = numerator.Mul(toMulWithNum)
			denominator = denominator.Mul(toMulWithNDen)
		}

		divRes := numerator.Div(denominator)

		ci = shares[i].Y.Mul(divRes)
		res = res.Add(ci) //, ci)
	}

	return res
}
