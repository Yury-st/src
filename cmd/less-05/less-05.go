package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]

	var numbers []float64

	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	if maximum(numbers...) != 0 {
		fmt.Println("Ваш максимальный аргумент: ", maximum(numbers...))

	}
}

func maximum(numbers ...float64) float64 {
	if len(numbers) == 0 {
		fmt.Println("Вы не ввели ни одого аргумента.")
		return 0
	}
	maxNum := numbers[0]
	for _, number := range numbers {
		if maxNum < number {
			maxNum = number
		}
	}
	return maxNum
}
