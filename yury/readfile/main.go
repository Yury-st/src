package main

import (
	"fmt"
	"os"
)

func openFile(fileName string) (*os.File, error) {
	fmt.Printf("File %v is opened\n", fileName)
	return os.Open(fileName)
}
