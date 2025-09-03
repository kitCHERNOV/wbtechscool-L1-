package l1_22

import (
	"fmt"
	"math/big"
)

type TwoBigNumbers struct {
	A *big.Int
	B *big.Int
}

func (tb *TwoBigNumbers) Multiply() big.Int {
	var result big.Int
	result.Mul(tb.A, tb.B)
	return result
}
func (tb *TwoBigNumbers) Addition() big.Int {
	var result big.Int
	result.Add(tb.A, tb.B)
	return result
}

func (tb *TwoBigNumbers) Substraction() big.Int {
	var result big.Int
	result.Sub(tb.A, tb.B)
	return result
}
func (tb *TwoBigNumbers) Division() big.Int {
	var result big.Int
	result.Div(tb.A, tb.B)
	return result
}

func MainFunc() {
	tbn := new(TwoBigNumbers)
	tbn.A = big.NewInt(100)
	tbn.B = big.NewInt(4)
	res := tbn.Multiply()
	fmt.Println("Операция умножения:", res.Int64())
	res = tbn.Division()
	fmt.Println("Операция деления:", res.Int64())
	res = tbn.Substraction()
	fmt.Println("Операция вычитания:", res.Int64())
	res = tbn.Addition()
	fmt.Println("Операция сложения:", res.Int64())
}
