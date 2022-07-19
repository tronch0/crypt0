package field

import (
	"fmt"
	"math/big"
)

type Element struct {
	order *big.Int
	n     *big.Int
}

func New(num, field *big.Int) *Element {
	return &Element{n: new(big.Int).Mod(num, field), order: field}
}

func (fe *Element) GetNum() *big.Int {
	return fe.n
}

func (fe *Element) GetOrder() *big.Int {
	return fe.order
}

func (fe *Element) Cmp(f *Element) int {
	checkEqualField(fe.order, f.order)

	return fe.n.Cmp(f.n)
}

func (fe *Element) Equal(f *Element) bool {
	checkEqualField(fe.order, f.order)

	return fe.n.Cmp(f.n) == 0
}

func (fe *Element) Add(f *Element) *Element {
	checkEqualField(fe.order, f.order)

	return &Element{
		n:     new(big.Int).Mod(new(big.Int).Add(fe.n, f.n), fe.order),
		order: new(big.Int).Set(fe.order),
	}
}

//func (fe *Element) AddInt(n *big.Int) *Element {
//	return &Element{
//		n:     new(big.Int).Mod(new(big.Int).Add(fe.n, n), fe.order),
//		order: new(big.Int).Set(fe.order),
//	}
//}

func (fe *Element) Sub(f *Element) *Element {
	checkEqualField(fe.order, f.order)

	res := new(big.Int).Sub(fe.n, f.n)

	if res.Cmp(big.NewInt(0)) < 0 {
		res = new(big.Int).Add(fe.order, res)
	}

	return &Element{
		n:     res,
		order: new(big.Int).Set(fe.order),
	}
}

func (fe *Element) Mul(f *Element) *Element {
	checkEqualField(fe.order, f.order)

	return &Element{
		n:     new(big.Int).Mod(new(big.Int).Mul(fe.n, f.n), fe.order),
		order: new(big.Int).Set(fe.order),
	}
}

func (fe *Element) ScalarMul(coef *big.Int) *Element {
	curr := fe.n
	res := big.NewInt(0)

	modCoef := new(big.Int).Mod(coef, fe.order)

	for modCoef.Cmp(big.NewInt(0)) > 0 {
		if modCoef.Bit(0) == 1 {
			res = new(big.Int).Add(res, curr)
		}
		curr = new(big.Int).Add(curr, curr)
		modCoef = new(big.Int).Rsh(modCoef, 1)
	}

	res = new(big.Int).Mod(res, fe.order)

	return &Element{
		n:     res,
		order: new(big.Int).Set(fe.order),
	}
}

func (fe *Element) Expo(k *big.Int) *Element {

	return &Element{
		n:     new(big.Int).Exp(fe.n, k, fe.order),
		order: new(big.Int).Set(fe.order),
	}
}

func (fe *Element) Div(f *Element) *Element {
	checkEqualField(fe.order, f.order)

	fnInv := new(big.Int).ModInverse(f.n, f.order)

	return &Element{
		n:     new(big.Int).Mod(new(big.Int).Mul(fe.n, fnInv), fe.order),
		order: new(big.Int).Set(fe.order),
	}
}

func Clone(p *Element) *Element {
	if p == nil {
		return nil
	}

	var newN *big.Int
	if p.n != nil {
		newN = new(big.Int).Set(p.n)
	}

	var newP *big.Int
	if p.order != nil {
		newP = new(big.Int).Set(p.order)
	}

	return &Element{order: newP, n: newN}
}

func (fe *Element) ToString() string {
	res := "{n: "
	if fe.n == nil {
		res += "nil, "
	} else {
		res += fe.n.String() + ", "
	}

	res += "order: "

	if fe.order == nil {
		res += "nil"
	} else {
		res += fe.order.String()
	}

	res += "}"

	return res
}

func checkEqualField(field1, field2 *big.Int) {
	if field1.Cmp(field2) != 0 {
		panic(fmt.Sprintf("cant compare field elements with diffrent fields (first-n-field: %d, second-n-field: %d)", field1, field2))
	}
}
