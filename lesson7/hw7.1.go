//Go. Уровень 2. Урок 7. Рефлексия и кодогенерация.
//Написать функцию, которая принимает на вход структуру in (struct или кастомную struct)
//и values map[string]interface{} (key - название поля структуры, которому нужно присвоить value этой мапы).
//Необходимо по значениям из мапы изменить входящую структуру in с помощью пакета reflect.
//Функция может возвращать только ошибку error.
//Написать к данной функции тесты (чем больше, тем лучше - зачтется в плюс).
//основная цель рефлексии - это определение типа и значения, хранящегося в объекте интерфейса
package main

import (
	"fmt"
	"go_level_2/go_level_2/lesson7/reflection"
	"log"
)

func main() {
	ValuesMap := map[string]interface{}{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	fmt.Printf("Begin struct Structure type %T %v\n", reflection.Structure, reflection.Structure)
	err := reflection.Reflection(ValuesMap, &reflection.Structure)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Finish struct Structure type %T %v\n", reflection.Structure, reflection.Structure)

}
