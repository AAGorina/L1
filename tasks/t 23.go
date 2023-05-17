package main

import "fmt"

/*Удалить i-ый элемент из слайса.*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var pos int = 7
	res := append(nums[0:pos], nums[pos+1:]...)
	fmt.Println(res)
}
