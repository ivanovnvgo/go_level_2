package scandirectory

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type statistic struct {
	name string
	size int64
}

var fileStructure statistic

//нужно решить: в чем хранить стат данные по всем файлам, в карте или слайсе. В карте данные перезаписываются, если ключ - имя файла.
var fileSlice = make([]statistic, 3)
var i int //счетчик файлов, индекс слайса

var filesMap = make(map[string]int64)

//ScanDirectory рекурсивная функция, которая получает путь для обработки всех файлов во всех вложенных директориях
func ScanDirectory(path string) error {
	fmt.Printf("current directory: %T %s\n", path, path) // печать текущей директории
	files, err := ioutil.ReadDir(path)                   //получаем слайс с содержимым текущей директории
	//fmt.Printf("slice with the contents of the current directory %T \n", files) //!!!!!!!!!!Удалить после отладки
	if err != nil {
		err = fmt.Errorf("error reading directory contents")
		return err
	}

	for _, file := range files {
		//fmt.Println(file.Name()) !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!Удалить после отладки
		//fmt.Printf("file: %T %s\n", file.Name(), file.Name()) //!!!!!!!!!!!!!!!!!!!!!!!!!!Удалить после отладки
		filePath := filepath.Join(path, file.Name()) //соединяем путь директории и имя файла через символ /
		if file.IsDir() {                            //если файл является директорией
			//рекурсивно вызываем функцию , в которую передаем путь вложенной директории (подкаталога)
			err := ScanDirectory(filePath)
			if err != nil {
				err = fmt.Errorf("subdirectory read error")
				return err
			}
		} else {
			fmt.Println(filePath) //если это не директория, а файл, то печатаем его имя и путь до него
			//fmt.Printf("filePath %T %s\n", filePath, filePath) //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!Удалить после отладки
			//чтение статистики файла
			fileStat, err := os.Stat(filePath)
			if err != nil {
				err = fmt.Errorf("error reading statistics")
				return err
			}
			//пишем в слайс
			fileStructure.name = fileStat.Name()
			fileStructure.size = fileStat.Size()
			if fileStat.Name() != "" && fileStat.Size() != 0 {
				fileSlice[i] = fileStructure
				i++
			}

			//пишем в карту
			filesMap[fileStat.Name()] = fileStat.Size()

			//fileStat: *os.fileStat &{scandirectory.go 2093 436 {788152527 63782346086 0x54aa00} {2053 141711 1 33204 1000 1000 0 0 2093 4096 8 {1646749291 516200475} {1646749286 788152527} {1646749286 788152527} [0 0 0]}}
			//fmt.Printf("type of name: %T, type of size: %T\n ", fileStat.Name(), fileStat.Size())//type of name: string, type of size: int64
			//fmt.Printf("fileStat: %T %v\n", fileStat, fileStat) //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!Удалить после отладки
			//fmt.Println(fileStructure) //{scandirectory.go 2764}
		}
	}
	//fmt.Println(fileSlice)
	//fmt.Println(filesMap)
	return nil
}
