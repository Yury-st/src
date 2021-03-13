package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func addTextToFile(text string, fileName string) {
	option := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile(fileName, option, os.FileMode(0600))
	check(err)
	defer file.Close()
	_, err = file.Write([]byte(text))
	check(err)
}

func main() {

	file, err := os.OpenFile("aardvark.txt", os.O_RDONLY, os.FileMode(0600))
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())
}
