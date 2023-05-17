package main

import (
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func ChangeBit(n int64, p int, v int) int64 {
	if v != 0 {
		// добоавить 1 - создаем маску 0000...010....000
		// при побитовой дьзьюнкции это даст 1 в нужной позиции
		mask := int64(1 << (int64(p) - 1))
		return n | mask
	} else {
		// создаем маску обратную случаю для 1
		mask := int64(-1 ^ 1<<(int64(p)-1))
		return n & mask
	}
}

func main() {
	var n int64 = 29
	var pose int = 1
	var v int = 0
	fmt.Println("Result: ", ChangeBit(n, pose, v))
}
