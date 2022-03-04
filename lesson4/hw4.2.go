//Написать программу, которая при получении в канал сигнала SIGTERM останавливается не позднее,
//чем за одну секунду (установить таймаут).
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		_, cancel := context.WithTimeout(context.Background(), 3*time.Second) //Освобождаем ресурсы
		defer cancel()
		//time.Sleep(1 * time.Second)//останавливается не позднее, чем за одну секунду (установить таймаут)
		fmt.Println("\nЗавершение программы")
		os.Exit(1)
	}()
	//Имитация зависания программы
	for {
		fmt.Println("Ожидание выполнения программы... для выхода нажмите ctrl+C")
		time.Sleep(60 * time.Second)
	}
}
