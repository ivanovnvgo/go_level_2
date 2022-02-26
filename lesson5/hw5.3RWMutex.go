//3.Протестируйте производительность операций чтения и записи на множестве действительных чисел,
//безопасность которого обеспечивается sync.RWMutex для разных вариантов использования:
//10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение
package main

import (
	"fmt"
	"sync"
	"time"
)

const maxRW = 100

//recordRW the function writes to a slice
func recordRW(s []int, x int) {
	for i := 0; i < maxRW-x; i++ {
		s[i] = i + 101
	}
}

//readRW the function reads from a slice
func readRW(s []int, x int) {
	for i := 0; i < maxRW-x; i++ {
		_ = s[i]
	}
}

func main() {
	var (
		mutexRW sync.RWMutex
	)
	sliceNumbers2 := make([]int, maxRW)
	for i := 0; i < maxRW; i++ {
		sliceNumbers2[i] = i + 1
	}

	fmt.Println("Use sync.RWMutex")

	startRWMutex := time.Now() // start time
	go func(s []int) {
		mutexRW.Lock()
		defer mutexRW.Unlock()
		recordRW(sliceNumbers2, 90)
	}(sliceNumbers2)
	go func(s []int) {
		mutexRW.RLock()
		defer mutexRW.RUnlock()
		readRW(sliceNumbers2, 10)
	}(sliceNumbers2)
	duration := time.Since(startRWMutex) // full time
	fmt.Println(duration)

	startRWMutex = time.Now() // start time
	go func(s []int) {
		mutexRW.Lock()
		defer mutexRW.Unlock()
		recordRW(sliceNumbers2, 50)
	}(sliceNumbers2)
	go func(s []int) {
		mutexRW.RLock()
		defer mutexRW.RUnlock()
		readRW(sliceNumbers2, 50)
	}(sliceNumbers2)
	duration = time.Since(startRWMutex) // full time
	fmt.Println(duration)

	startRWMutex = time.Now() // start time
	go func(s []int) {
		mutexRW.Lock()
		defer mutexRW.Unlock()
		recordRW(sliceNumbers2, 10)
	}(sliceNumbers2)
	go func(s []int) {
		mutexRW.RLock()
		defer mutexRW.RUnlock()
		readRW(sliceNumbers2, 90)
	}(sliceNumbers2)
	duration = time.Since(startRWMutex) // full time
	fmt.Println(duration)
}
