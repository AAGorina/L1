package main

import "fmt"

/* Реализовать пересечение двух неупорядоченных множеств. */

func intersection(x []int, y []int) []int {
	var result []int
	cheker := make(map[int]int)
	for _, it := range x {
		cheker[it] = 1
	}
	for _, it := range y {
		_, ok := cheker[it]
		if ok {
			result = append(result, it)
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	fmt.Println(intersection(a, b))
}
