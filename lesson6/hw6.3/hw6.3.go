//Go. Уровень 2. Домашнее задание к уроку 6
//6.3.Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”
//$ while true; do go run -race hw6.3.go; sleep 1; done
//or
//$ go run -race hw6.3.go
package main

import (
	"fmt"
)

const max = 100

//record the function writes to a slice
func record(s []int, x int) {
	for i := 0; i < max-x; i++ {
		s[i] = i + 101
	}
}

//read the function reads from a slice
func read(s []int, x int) {
	for i := 0; i < max-x; i++ {
		_ = s[i]
	}
}

func main() {

	var (
	//mutex sync.Mutex
	)
	sliceNumbers1 := make([]int, max)

	for i := 0; i < max; i++ {
		sliceNumbers1[i] = i + 1
	}

	fmt.Println("start")
	for i := 0; i < 10000; i++ {
		go func(s []int) {
			//mutex.Lock()
			//defer mutex.Unlock()
			record(sliceNumbers1, 0)
		}(sliceNumbers1)
		go func(s []int) {
			//mutex.Lock()
			//defer mutex.Unlock()
			read(sliceNumbers1, 10)
		}(sliceNumbers1)

		go func(s []int) {
			//mutex.Lock()
			//defer mutex.Unlock()
			record(sliceNumbers1, 0)
		}(sliceNumbers1)
		go func(s []int) {
			//mutex.Lock()
			//defer mutex.Unlock()
			read(sliceNumbers1, 90)
		}(sliceNumbers1)

		record(sliceNumbers1, 0)
		read(sliceNumbers1, 50)

		record(sliceNumbers1, 0)
		read(sliceNumbers1, 90)
	}
	fmt.Println("finish")
}
