package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func printSpace(path string) string {
	spaces := ""
	for i := 0; i < len(path); i++ {
		spaces += " "
	}
	return spaces + "|_"
}

func scanDirectory(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			fmt.Printf("%s%s\n", printSpace(path), file.Name())
			scanDirectory(filePath)
		} else {
			fmt.Printf("%s%s\n", printSpace(path), file.Name())
		}
	}
	return nil
}

func main() {
	path := ".."
	fmt.Println("|__")
	err := scanDirectory(path)
	if err != nil {
		log.Fatal(err)
	}
}
