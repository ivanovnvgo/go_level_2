package reflection

import (
	"fmt"
	"reflect"
)

type In struct {
	Key int
}

var Structure In

// Reflection функция заполняет поля структуры значениями из карты, возвращает только ошибку error.
func Reflection(m map[string]interface{}, Structure In) error {
	//fmt.Println(reflect.TypeOf(m))
	for k, v := range m {
		//fmt.Println("Key:", k, "Value:", v)
		//fmt.Println(reflect.TypeOf(k))
		Structure.Key = int(reflect.ValueOf(v).Int())
		//i.Key = v.(int) - это работает, но не использует reflect
		fmt.Println(Structure.Key, Structure)

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
