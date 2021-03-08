package main

import "fmt"

func main() {
	defer fmt.Println("       -----function main")
	one()
}

func one() {
	defer fmt.Println("       -----function one")
	two()
}
func two() {
	defer fmt.Println("       -----function two")
	three()
}
func re() {
	recover()
}
func three() {
	defer fmt.Println("       -----function three")
	defer re()
	panic("Hello))) I'm a panic.")
}
