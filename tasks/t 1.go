package main

import "fmt"

/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human
(аналог наследования).*/

// структура Human
type Human struct {
	Name       string
	Age        int
	PasportNum string
}

// методы структуры Human
func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) GetAge() int {
	return h.Age
}

func (h *Human) GetPasportNum() string {
	return h.PasportNum
}

// структура Action
type Action struct {
	Human
}

// перепишем метод GetPasportNum
func (a *Action) GetPasportNum() string {
	return "Access denied!"
}

func main() {
	h := Human{
		Name:       "It's a name",
		Age:        23,
		PasportNum: "pretend it's a pasport",
	}

	a := Action{
		h,
	}

	// проверка методов Human
	fmt.Println(h.GetName())
	fmt.Println(h.GetAge())
	fmt.Println(h.GetPasportNum())
	// проверка методов Action
	fmt.Println(a.GetName())
	fmt.Println(a.GetAge())
	// методы наследуются
	// метод изменился
	fmt.Println(a.GetPasportNum())
}
