package main

import "fmt"

/*Реализовать паттерн «адаптер» на любом примере.*/
/* адаптер - если интерфес не совместим с системой вместо переписывания методов делаем адаптер*/

type lightning interface {
	is_on() bool
}

type check_light struct {
}

func (c *check_light) Light(l lightning) bool {
	if l.is_on() {
		return true
	}
	return false
}

type bulb struct {
	brightness int
}

func (b *bulb) is_on() int {
	return b.brightness
}

type adapt_bulb struct {
	b bulb
}

func (a *adapt_bulb) is_on() bool {
	return a.b.brightness >= 500
}

func main() {
	b := adapt_bulb{
		bulb{520},
	}
	c_light := check_light{}
	fmt.Println(c_light.Light(&b))
}
