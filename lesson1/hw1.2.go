//Первое домашнее задание по курсу "Go. Уровень 2"
//Урок 1. Продвинутая генерация и обработка ошибок и аварий (паник). Отложенный вызов функций
//1.1. Для закрепления навыков отложенного вызова функций, напишите программу, содержащую вызов функции,
//которая будет создавать паническую ситуацию неявно.
//Затем создайте отложенный вызов, который будет обрабатывать эту паническую ситуацию и, в частности,
//печатать предупреждение в консоль.
//Критерием успешного выполнения задания является то, что программа не завершается аварийно ни при каких условиях.
//1.2. Дополните функцию из п.1 возвратом собственной ошибки в случае возникновения панической ситуации.
//Собственная ошибка должна хранить время обнаружения панической ситуации.
//Критерием успешного выполнения задания является наличие обработки созданной ошибки в функции main
//и вывод ее состояния в консоль
package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

//ErrorWithTrace - создаю свою ошибку как это показано в методичке
//type ErrorWithTrace struct {
//	text  string
//	trace string
//}

//New возвращает ошибку в виде структуры с двумя полями
//func New(text string) error {
//	return &ErrorWithTrace{
//		text:  text,
//		trace: string(debug.Stack()),
//	}
//}

//Error - метод, который формирует  и возвращает одну строку, содержащую текст ошибки и ее стектрейс
//func (e *ErrorWithTrace) Error() string {
//	return fmt.Sprintf("error: %s\n trace\n%s", e.text, e.trace)
//}

//calculator2 - моя функция с примером формирования собственной ошибки
func calculator2() (float64, error) {
	var a, b float64
	var op string
	fmt.Println("Введите первое число, оператор (*, /, +, -, %(Процент), второе число. После ввода каждого значения нажимайте клавишу ENTER: ")
	fmt.Scan(&a, &op, &b)
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "%":
		return a * b / 100, nil
	case "/":
		//if b == 0 {
		//	panic("second operand is ZERO!!!")
		//	//return 0, fmt.Errorf("second operand is zero %w", err)
		//
		//} else {
		//	return a / b, nil
		//}
		return a / b, errors.New(fmt.Sprintln("second operand is zero, time of panic: ", time.Now())) //создал собственную ошибку
	default:
		return 0, errors.New(fmt.Sprintln("operation is not found , time of error: ", time.Now())) //создал собственную ошибку
	}
}

func main() {

	defer func() {
		v := recover()
		err, ok := v.(error)
		if !ok {
			fmt.Println("ok in the recover is", ok)
		}
		if err != nil {
			//err = errors.Wrap(err, "recover() caught a panic")
			//Получил при запуске и я не понял, почему это произошло: ./hw1.2.go undefined: errors.Wrap
			log.Printf("recovered %s", err)
		}
		log.Println("recovered", v) //Собственная ошибка хранит время обнаружения панической ситуации.
	}()

	//defer func() {
	//	if v := recover(); v != nil {
	//		fmt.Println("recovered", v)
	//	}
	//}()
	result, err := calculator2()
	if err != nil {
		//Критерием успешного выполнения задания является наличие обработки созданной ошибки в функции main
		//и вывод ее состояния в консоль
		//Собственная ошибка хранит время обнаружения панической ситуации
		log.Printf("Error Encountered, err %w: ", err)
	}
	fmt.Println("result= ", result)
}
