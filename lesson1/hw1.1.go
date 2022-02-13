//Первое домашнее задание по курсу "Go. Уровень 2"
//Урок 1. Продвинутая генерация и обработка ошибок и аварий (паник). Отложенный вызов функций
//1.1. Для закрепления навыков отложенного вызова функций, напишите программу, содержащую вызов функции,
//которая будет создавать паническую ситуацию неявно.
//Затем создайте отложенный вызов, который будет обрабатывать эту паническую ситуацию и, в частности,
//печатать предупреждение в консоль.
//Критерием успешного выполнения задания является то, что программа не завершается аварийно ни при каких условиях.

//Решение:
//1.Напишите программу, содержащую вызов функции, которая будет создавать паническую ситуацию неявно.
//Вызываем функцию calculator1
package main

import "fmt"

func calculator1() (float64, error) {
	var a, b float64
	var op string
	var err error
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
		if b == 0 {
			return a / b, fmt.Errorf("no panic => second operand is ZERO!!! %w", err)
			//panic("this is a panic => second operand is ZERO!!!") //Я пытался создать неявную панику,
			//но компилятор деление на 0 пропустил без ошибки, значение +Inf
		} else {
			return a / b, nil
		}
	default:
		//return 0, fmt.Errorf("operation is not found %w", err)
		//Для вызова неявной паники обращаемся к несуществующему элементу слайса
		slice := []float64{0, 1, 2, 3}
		return slice[4], nil // эту панику поймал recover()
	}

}

func main() {
	//2.Затем создайте отложенный вызов, который будет обрабатывать эту паническую ситуацию и, в частности,
	//печатать предупреждение в консоль.
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
		fmt.Println("recover end!")
	}()
	result, err := calculator1()
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("result= ", result)
	fmt.Println("main end!")
}
