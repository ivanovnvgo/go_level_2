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
	"log"
	"reflect"
)

//type appropriator interface {
//	reflectionExample(m map[string]interface{}) error
//}

type In struct {
	Key int
}

//ReflectionExample функция возвращает только ошибку error.
func (i *In) Reflection(m map[string]interface{}) error {
	//fmt.Println(reflect.TypeOf(m))
	for k, v := range m {
		//fmt.Println("Key:", k, "Value:", v)
		//fmt.Println(reflect.TypeOf(k))
		i.Key = int(reflect.ValueOf(v).Int())
		//i.Key = v.(int) - это работает, но не использует reflect
		fmt.Println(i.Key, i)

		t := reflect.ValueOf(k)                       //получаем значение ключа карты
		fmt.Println(t, "have type of key:", t.Type()) //смотрим тип значения ключа карты
		keyMap := reflect.TypeOf(k)                   //получаем тип значения ключа карты
		fmt.Println("type key of map= ", keyMap)
		/*
			//проверка соответствия типа ключа карты (мапы)
			//witch работает только со значением типа interface
			//switch t.(type) { //cannot type switch on non-interface value t (type reflect.Value)
			switch k.(type) { //cannot type switch on non-interface value k (type string)
			case string:
				return nil
			case int:
				return fmt.Errorf("type mismatch, expected value as string")
			case float64:
				return fmt.Errorf("type mismatch, expected value as string")
			case bool:
				return fmt.Errorf("type mismatch, expected value as string")
			}
		*/

		//проверка соответствия типа значения элемента карты (мапы)
		switch v.(type) {
		case string:
			return fmt.Errorf("type mismatch, expected value as int")
		case int:
			return nil
		case float64:
			return fmt.Errorf("type mismatch, expected value as int")
		case bool:
			return fmt.Errorf("type mismatch, expected value as int")
		}
	}
	return nil
}

func main() {
	var structure In

	valuesMap := map[string]interface{}{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	fmt.Printf("Begin struct structure %v\n", structure)
	err := structure.Reflection(valuesMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Finish struct structure %v\n", structure)

}
