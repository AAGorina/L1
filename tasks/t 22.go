package main

import (
	"fmt"
	"math/big"
)

/*Разработать программу, которая
перемножает, делит, складывает, вычитает
две числовых переменных a,b, значение которых > 2^20.*/

func main() {
	var a big.Float
	var b big.Float
	var c big.Float
	a.SetString("85864269754258908754")
	b.SetString("79488268387823487298")
	fmt.Println(c.Add(&a, &b)) // сложение
	fmt.Println(c.Sub(&a, &b)) //вычитание
	fmt.Println(c.Mul(&a, &b)) //умножение
	fmt.Println(c.Quo(&a, &b)) //деление
}
