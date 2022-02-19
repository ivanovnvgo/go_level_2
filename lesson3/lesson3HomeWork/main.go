//hw2.go выводит на печать строку, возвращаемую функцией HelloWorld()
package main

import (
	"fmt"
	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
)

//enterName запрашивает ввод имени
func enterName() string {
	var s string
	fmt.Print("Введите Ваше имя: ")
	fmt.Scanf("%s\n", &s)
	return s
}

//HelloWorld форматирует возвращаемую функцией строку
func HelloWorld() string {
	var name string = enterName()
	return fmt.Sprintf("Hello, %s!!!", name)
}
func main() {

	fmt.Println(HelloWorld())
}
