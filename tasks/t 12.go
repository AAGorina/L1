package main

import "fmt"

/*Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество.*/

func main() {
	smth := []string{"cat", "cat", "dog", "cat", "tree"}
	uniq := make(map[string]int)
	var result []string
	for _, it := range smth {
		_, ok := uniq[it]
		if !ok {
			uniq[it] = 1
			result = append(result, it)
		}
	}
	fmt.Println(result)
}
