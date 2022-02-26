//1.Напишите программу, которая запускает n потоков и дожидается завершения их всех
//2.Реализуйте функцию для разблокировки мьютекса с помощью defer

package main

import (
	"fmt"
	"sync"
	"time"
)

//counter - counting the number of running goroutines
//func counter(count *int) *int {
//	*count++
//	return count
//}

func main() {
	const n = 5
	var (
		mutex sync.Mutex
		ch    = make(chan int, n)
		wg    = sync.WaitGroup{} //создаем экземпляр семафора для реализации ожидания завершения всех горутин
	)
	count := 0
	wg.Add(n) //инициализируем семафор исходным состоянием (не дает задать число, отличное от n=5)
	for i := 0; i < n; i++ {
		//1.Напишите программу, которая запускает n потоков и дожидается завершения их всех
		go func(i int) {
			ch <- i
			// Захват мьютекса
			mutex.Lock()
			//2.Реализуйте функцию для разблокировки мьютекса с помощью defer
			// Освобождение мьютекса
			defer mutex.Unlock()
			//counter(&count) //count++ //counting the number of running goroutines
			func() {
				count++
				wg.Done() //выполняем декремент семафора
			}()
			fmt.Printf("i'm %d-goroutine, count= %d\n", i, count)
		}(i)
	}
	//не могу убрать эту задержку, без нее код завершается с ошибкой:
	//all 0 goroutines are completed
	//panic: send on closed channel
	time.Sleep(1 * time.Second) //
	close(ch)
	//Что-то делаем:
	i := 0
	for range ch {
		i += 1 //counting the number of records in the channel ch
	}
	if i == count {
		fmt.Printf("all %d goroutines are completed\n", i)
	}
	wg.Wait() //ждем обнуления семафора
}
