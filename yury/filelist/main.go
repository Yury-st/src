//Этот код позволяет получить список файлов в директории
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func formatedPath(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			// fmt.Println(path[:int(i)])
			return strings.Repeat(" ", int(i)) + "|_" + path[int(i+1):]
		}
	}
	return path
}

//Функция checkList получает название папки и отдает слайс
//списка файлов и папок в этой папке
func scanDirectory(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())

		if file.IsDir() {
			fmt.Println(formatedPath(filePath))
			err := scanDirectory(filePath)

			if err != nil {
				return err
			}
		} else {
			fmt.Println(formatedPath(filePath))
		}

	}
	return nil
}

func main() {
	FolderName := "/Users/yury/go/src"
	err := scanDirectory(FolderName)
	if err != nil {
		log.Fatal(err)
	}
}
