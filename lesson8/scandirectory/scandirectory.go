// Package scandirectory принимает на вход флаг с путем директории для сканирования и флаг на удаление дублирующихся файлов.
// выводит список дублирующихся файлов
package scandirectory

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type statistic struct {
	name string
	size int64
}

var (
	filesMap = make(map[int]string)
	j        int
)

// ScanDirectory рекурсивная функция, которая получает путь для обработки всех файлов во всех вложенных директориях
func ScanDirectory(path string) error {
	//fmt.Printf("current directory: %T %s\n", path, path) // печать текущей директории
	files, err := ioutil.ReadDir(path) //получаем слайс с содержимым текущей директории
	if err != nil {
		err = fmt.Errorf("error reading directory contents")
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name()) //соединяем путь директории и имя файла через символ /
		if file.IsDir() {                            //если файл является директорией
			//рекурсивно вызываем функцию , в которую передаем путь вложенной директории (подкаталога)
			err := ScanDirectory(filePath)
			if err != nil {
				err = fmt.Errorf("subdirectory read error")
				return err
			}
		} else {
			//fmt.Println(filePath) //если это не директория, а файл, то печатаем его имя и путь до него
			//пишем в карту
			filesMap[j] = filePath
			j++
		}
	}
	return nil
}

// PrintNameDoubleFiles находит в карте дублирующиеся файлы и выводит их на печать
func PrintNameDoubleFiles() error {
	for i := 0; i < len(filesMap); i++ {
		a, err := os.Stat(filesMap[i])
		if err != nil {
			err = fmt.Errorf("error reading statistics")
			return err
		}
		for j := i + 1; j < len(filesMap); j++ {
			b, err := os.Stat(filesMap[j])
			if err != nil {
				err = fmt.Errorf("error reading statistics")
				return err
			}
			if a.Name() == b.Name() && a.Size() == b.Size() {
				fmt.Println("duplicate files:")
				fmt.Printf("%s = %s\n", filesMap[j], filesMap[i])
			}
		}
	}
	return nil
}

// ReadFlagDeleteDoubleFile принимает от пользователя флаги:
// путь к нужной директории и возможность удаления повторяющегося файла
func ReadFlagDeleteDoubleFile() (string, string) {
	var pathDir string
	var rmFile string
	//var pathDirectory = os.Args[1:2] //запись аргумента командной строки - читаем путь к нужной директории
	//for i := 0; i < len(pathDirectory); i++ {
	//	pathDir = pathDir + pathDirectory[i]
	//}
	var pFlag = flag.String("pFlag", "/home/user/go/src/go_level_2/go_level_2/lesson8/scandirectory", "путь к нужной директории")
	var rFlag = flag.String("rFlag", "n", "флаг удаления повторяющегося файла")
	flag.Parse()
	pathDir = *pFlag
	rmFile = *rFlag
	fmt.Printf("флаг удаления повторяющегося файла установлен: %s\n", rmFile)
	fmt.Printf("сканируем директорию %s\n", pathDir)
	return pathDir, rmFile
}
