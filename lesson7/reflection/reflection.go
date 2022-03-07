package reflection

import (
	"fmt"
	"reflect"
)

type In struct {
	Key int
}

var Structure In

// Reflection функция заполняет поля структуры случайными значениями из карты, возвращает только ошибку error.
func Reflection(m map[string]interface{}, structure *In) error {
	//fmt.Println(reflect.TypeOf(m))
	for k, v := range m {
		//fmt.Println("Key:", k, "Value:", v)
		//fmt.Println(reflect.TypeOf(k))
		Structure.Key = int(reflect.ValueOf(v).Int())
		//i.Key = v.(int) - это работает, но не использует reflect
		fmt.Printf("structure have type %T and have value = %v\n ", Structure, Structure.Key)

		t := reflect.ValueOf(k)                       //получаем значение ключа карты
		fmt.Println(t, "have type of key:", t.Type()) //смотрим тип значения ключа карты
		keyMap := reflect.TypeOf(k)                   //получаем тип значения ключа карты
		fmt.Println("type key of map= ", keyMap)

		//проверка соответствия типа значения элемента карты (мапы)
		value := reflect.ValueOf(v)
		typeValue := reflect.TypeOf(v)
		fmt.Printf("value %T, type %v\n ", value, typeValue)
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
