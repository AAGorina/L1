package main

import (
	"fmt"
	"strings"
)

/*Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.*/

func CheckUnique(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]int)
	for _, r := range str {
		_, ok := m[r]
		if ok {
			return false
		}
		m[r] = 1
	}
	return true
}

func main() {
	fmt.Println(CheckUnique("abCdefAaf"))
}
