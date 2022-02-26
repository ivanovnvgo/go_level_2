//3.Протестируйте производительность операций чтения и записи на множестве действительных чисел,
//безопасность которого обеспечивается sync.Mutex для разных вариантов использования:
//10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение
package main

import (
	"fmt"
	"sync"
	"time"
)

const max = 100

//record the function writes to a slice
func record(s []int, x int) {
	for i := 0; i < max-x; i++ {
		s[i] = i + 1
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
		mutex sync.Mutex
	)
	sliceNumbers1 := make([]int, max)
	for i := 0; i < max; i++ {
		sliceNumbers1[i] = i + 1
	}

	fmt.Println("Use sync.Mutex")

	startMutex := time.Now() // start time
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
	duration := time.Since(startMutex) // full time
	fmt.Println(duration)

	startMutex = time.Now() // start time
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
	duration = time.Since(startMutex) // full time
	fmt.Println(duration)

	startMutex = time.Now() // start time
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
	duration = time.Since(startMutex) // full time
	fmt.Println(duration)
}
