package main

import "fmt"

/*Поменять местами два числа без создания временной переменной.*/

func main() {
	nums := []int{1, 2}
	nums[0], nums[1] = nums[1], nums[0]
	fmt.Println(nums)
}
