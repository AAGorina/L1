package main

import (
	"fmt"
	"time"
)

/*Реализовать собственную функцию sleep.*/

func sleep(t time.Duration) {
	to := time.Now().Add(t)
	for time.Now().Unix() < to.Unix() {
	}
}

func main() {
	sleepfor := 3 * time.Second
	fmt.Println("Start")
	sleep(sleepfor)
	fmt.Println("End")
}
