package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

/* давайте посчитаем суммарное число каждой циферке в данных
две горутины + изолируем словарь при добавлении данных*/

func count_1_2_3(dat []int, wg *sync.WaitGroup, count map[int]int, m *sync.Mutex) {
	defer wg.Done()
	for _, item := range dat {
		m.Lock()
		_, ok := count[item]
		if ok {
			count[item] = count[item] + 1
		} else {
			count[item] = 1
		}
		m.Unlock()
	}
}

func main() {
	numbers := make(map[int]int)
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	data := []int{1, 2, 2, 3, 1, 2, 3, 2, 3, 2, 1, 2, 1, 2, 3, 2, 1, 2, 3,
		1, 2, 3, 1, 2, 3, 1, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 2, 3, 1,
		2, 3, 1, 3, 1, 2, 1, 2, 3, 2, 1, 1, 1, 2, 3, 3, 2, 2, 3, 3, 3, 1}
	data_1 := []int{2, 3, 2, 2, 3, 1, 1, 3, 2, 3, 1, 2, 2, 3, 3, 1, 1, 2,
		3, 2, 2, 1, 2, 3, 1, 2, 1, 3, 2, 1, 3, 2, 3, 3, 2, 3, 1, 1, 3, 1,
		2, 3, 3, 1, 1, 1, 1, 2, 3, 2, 1, 2, 3, 1, 1, 2, 1, 3, 2, 3, 1, 2}
	wg.Add(2)
	go count_1_2_3(data, &wg, numbers, &m)
	go count_1_2_3(data_1, &wg, numbers, &m)
	wg.Wait()
	fmt.Println(numbers)
}
