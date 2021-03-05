//Этот код позволяет получить список файлов в директории
package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
)

func checkList(folder string) []fs.FileInfo {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func returnList(files []fs.FileInfo) {
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory: ", file.Name())
			returnList(checkList(file.Name()))
		} else {
			fmt.Println("File: ", file.Name())
		}
	}
}

func main() {
	files := checkList("yury")
	returnList(files)

}
