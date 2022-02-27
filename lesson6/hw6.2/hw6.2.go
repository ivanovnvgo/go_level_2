//Go. Уровень 2. Домашнее задание к уроку 6
//6.2.Написать многопоточную программу, в которой будет использоваться явный вызов планировщика.
//"Выполнить подобную операцию планировщик может в случае завершения системного вызова, завершения синхронизации
//или атомарной операции и даже по истечении кванта времени."
//Выполните трассировку программы
//$ GOMAXPROCS=1 go run hw6.2.go 2>trace.out
//$ go tool trace trace.out
//перейти по сформировавшейся ссылке в терминале в браузер

package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"sync/atomic"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	trace.Start(os.Stderr) //Выполните трассировку программы
	defer trace.Stop()
	wg.Add(3)
	go func() {
		var (
			a = int64(2)
			r = atomic.SwapInt64(&a, 1) //атомарная операция
		)
		fmt.Println(a, r)
		wg.Done()
	}()
	go func() {
		var a = int64(2)
		atomic.CompareAndSwapInt64(&a, 2, 1) //атомарная операция
		fmt.Println(a)
		wg.Done()
	}()
	go func() {
		var (
			a = int64(2)
			r = atomic.AddInt64(&a, 1) //атомарная операция
		)
		fmt.Println(a, r)
		wg.Done()
	}()
	wg.Wait()
}
