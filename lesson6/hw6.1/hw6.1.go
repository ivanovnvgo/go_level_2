//Go. Уровень 2. Домашнее задание к уроку 6
//6.1.Написать программу, которая использует мьютекс для безопасного доступа к данным из нескольких потоков.
//Выполните трассировку программы
//$ GOMAXPROCS=1 go run hw6.1.go 2>trace.out
//$ go tool trace trace.out
//перейти по сформировавшейся ссылке в терминале в браузер

package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
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
	trace.Start(os.Stderr) //Выполните трассировку программы
	defer trace.Stop()
	var (
		mutex sync.Mutex
	)
	sliceNumbers1 := make([]int, max)

	for i := 0; i < max; i++ {
		sliceNumbers1[i] = i + 1
	}

	fmt.Println("Use sync.Mutex")

	go func(s []int) {
		mutex.Lock()
		defer mutex.Unlock()
		record(sliceNumbers1, 90)
	}(sliceNumbers1)
	go func(s []int) {
		mutex.Lock()
		defer mutex.Unlock()
		read(sliceNumbers1, 10)
	}(sliceNumbers1)

	go func(s []int) {
		mutex.Lock()
		defer mutex.Unlock()
		record(sliceNumbers1, 50)
	}(sliceNumbers1)
	go func(s []int) {
		mutex.Lock()
		defer mutex.Unlock()
		read(sliceNumbers1, 50)
	}(sliceNumbers1)

	go func(s []int) {
		mutex.Lock()
		defer mutex.Unlock()
		record(sliceNumbers1, 10)
	}(sliceNumbers1)
	go func(s []int) {
		mutex.Lock()
		defer mutex.Unlock()
		read(sliceNumbers1, 90)
	}(sliceNumbers1)
}
