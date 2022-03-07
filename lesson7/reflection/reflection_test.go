//Unit.test, запуск: go test reflection_test.go
package reflection

import (
	reflection "go_level_2/go_level_2/lesson7/reflection"
	"log"
	"reflect"
	"testing"
)

// Проверка на равенство с эталонной последовательностью: табличный тест
func TestReflectionTable(t *testing.T) {
	valuesMap := map[string]interface{}{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	reflection.Reflection(valuesMap, &reflection.Structure)
	s := reflect.ValueOf(&reflection.Structure).Elem()
	count := 0
	for i := 0; i < s.NumField(); i++ {
		for _, v := range valuesMap {
			//fmt.Println(v, reflection.Structure.Key)
			if v == reflection.Structure.Key {
				count++
			}
		}
	}
	if count == 0 {
		log.Fatalf("присвоены не табличные значения карты")
	}
}

// Проверка на ввод не числа int: тест на задымление
func TestReflectionSmoke(t *testing.T) {
	valuesMap := map[string]interface{}{
		"rsc": 3711.09,
		"r":   "2138",
		"gri": true,
	}
	var x = reflection.Reflection(valuesMap, &reflection.Structure)
	if x == nil {
		t.Fatalf("Было введено число не int, а ошибка не получена")
	}
}
