package main

import (
	"bufio"
	"fmt"
	"os"
)

/*Разработать программу, которая переворачивает подаваемую
на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.*/

func Reverce(s string) string {
	r := []rune(s)
	for i := 0; i*2 < len(r); i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}
	return string(r)
}

func main() {
	var str string
	in := bufio.NewReader(os.Stdin)
	str, _ = in.ReadString('\n')
	fmt.Println(Reverce(str))
}
