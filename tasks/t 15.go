package main

import (
	"fmt"
	"strings"
)

/*К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

Проблема - слайс это ссылочный тип. Мы забьем память если будем часто использовать эту функцию тк
память на строку будет выделяться, но не всей ей мы пользуемся в итоге
*/

/*
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//b := make([]byte, 100)
	// никуда не передаем слайс, память можно очистить после окончания работы функции
	//copy(b, v[:100])
	// или через строки
	justString = strings.Clone(v[:100])
	//justString = string(b)
}

func createHugeString(size int) string {
	// встроеннный метод работы сос троками для минимизации копирования памяти
	var sb strings.Builder
	for i := 0; i < size; i++ {
		// обавление символа
		sb.WriteString("s")
	}
	return sb.String()
}

func main() {
	someFunc()
	fmt.Println(justString)
}
