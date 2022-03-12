// Package scandirectory_test формирует примеры использования функций
package scandirectory_test

import (
	"go_level_2/go_level_2/lesson8/scandirectory"
	"log"
)

func ExampleReadFlagDeleteDoubleFile() {
	path, rm := scandirectory.ReadFlagDeleteDoubleFile()
	_, _ = path, rm
	// Output:флаг удаления повторяющегося файла установлен: n
	//сканируем директорию /home/user/go/src/go_level_2/go_level_2/lesson8/scandirectory
}
func ExampleScanDirectory() {
	path := "/home/user/go/src/go_level_2/go_level_2/lesson8/scandirectory"
	err := scandirectory.ScanDirectory(path)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
}
func ExamplePrintNameDoubleFiles() {
	err := scandirectory.PrintNameDoubleFiles()
	if err != nil {
		log.Fatal(err)
	}
	// Output:duplicate files:
	///home/user/go/src/go_level_2/go_level_2/lesson8/scandirectory/text.txt = /home/user/go/src/go_level_2/go_level_2/lesson8/scandirectory/subscandirectory/text.txt

}
