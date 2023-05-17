package main

import (
	"fmt"
	"reflect"
)

/*Разработать программу, которая в рантайме
способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.*/

// %T в пакете fmt
func GetTypeFMT(mistery interface{}) string {
	return fmt.Sprintf("%T", mistery)
}

// использовать переключатель
func GetTypeSwitch(mistery interface{}) string {
	var result string
	switch mistery.(type) {
	case int:
		result = "integer"
	case string:
		result = "string"
	case bool:
		result = "bool"
	default:
		result = "unknown"
	}
	return result
}

// через reflect
func GetTypeReflect(mistery interface{}) reflect.Type {
	return reflect.TypeOf(mistery)
}

func main() {
	ch := make(chan int)
	fmt.Println(GetTypeFMT(23.6))
	fmt.Println(GetTypeSwitch(false))
	fmt.Println(GetTypeReflect(ch))
	close(ch)
}
