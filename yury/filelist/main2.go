package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func lineToSpace(line string) string {
	lineSpaces := ""
	for i := 0; i < len(line); i++ {
		lineSpaces += " "
	}
	return lineSpaces + "|_"
}

func scanDirectory(path string) error {

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	} else {
		for _, file := range files {
			filePath := filepath.Join(path, file.Name())
			length := len(filePath) - len(file.Name())
			if file.IsDir() {
				fmt.Printf("%s%s\n", lineToSpace(path), filePath[length:])
				err := scanDirectory(filePath)
				if err != nil {
					return err
				}

			} else {
				fmt.Printf("%s%s\n", lineToSpace(path), filePath[length:])

			}
		}
	}
	return nil
}

func main() {
	path := "../"
	fmt.Printf("%s\n", "|__")
	err := scanDirectory(path)
	if err != nil {
		log.Fatal(err)
	}

}
