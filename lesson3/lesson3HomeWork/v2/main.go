//hw2.go выводит на печать строку, возвращаемую функцией HelloWorld()
package main

import (
	"fmt"
	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
	"os"
)

//enterName запрашивает ввод имени
func enterName() []string {
	//var s string
	//fmt.Print("Введите Ваше имя: ")
	//fmt.Scanf("%s\n", &s)
	//return s
	fmt.Print("Ваше имя - ")
	s := os.Args[1:] //Запись аргументов командной строки
	// (исключая нулевой элемент - название самой программы) в слайс
	fmt.Println(s)
	return s
}

//HelloWorld форматирует возвращаемую функцией строку
func HelloWorld() string {
	var name []string = enterName()
	return fmt.Sprintf("Hello, %s!!!", name)
}
func main() {

	fmt.Println(HelloWorld())
}
